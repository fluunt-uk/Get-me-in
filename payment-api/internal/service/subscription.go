package service

import (
	"encoding/json"
	"github.com/ProjectReferral/Get-me-in/payment-api/configs"
	"github.com/ProjectReferral/Get-me-in/payment-api/internal/models"
	"github.com/ProjectReferral/Get-me-in/payment-api/lib/rabbitmq"
	stripe_api "github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api"
	"github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api/resources/card"
	customer "github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api/resources/customer"
	sub "github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api/resources/subscription"
	token "github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api/resources/token"
	"github.com/ProjectReferral/Get-me-in/pkg/http_lib"
	"github.com/stripe/stripe-go"
	"net/http"
	"sync"
)

type Subscription struct {
	wg *sync.WaitGroup
	sync.RWMutex
	CustomerClient customer.Builder
	SubClient sub.Builder
	TokenClient token.Builder
	CardClient card.Builder
}

func (s *Subscription) Init(){
	s.wg = &sync.WaitGroup{}
}

func (s *Subscription) SubscribeToPremiumPlan(w http.ResponseWriter, r *http.Request){

	body := &models.Subscriber{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if !stripe_api.HandleError(err, w) {

		c := &stripe.Customer{}
		t := &stripe.Token{}
		var err error
		//create new customer and token
		s.asyncRequest(body, c, t, err)
		s.wg.Wait()

		if !stripe_api.HandleError(err, w) {

			//link the card with customer
			_, cardErr := s.CardClient.Put(c, t)

			if !stripe_api.HandleError(cardErr, w) {

				//create the sub and make payment
				sm, subErr := s.SubClient.Put(c, configs.PREMIUM_PLAN)

				b, _ := json.Marshal(models.ChangeRequest{
					Field:   "premium",
					NewBool: true,
					Type:    3,
				})

				//change premium field to true on user's account
				_, err := http_lib.Patch(configs.ACCOUNT_API_PREMIUM, b,
					map[string]string{configs.AUTH_HEADER: r.Header.Get(configs.AUTH_HEADER)})

				if !stripe_api.HandleError(subErr, w) && !stripe_api.HandleError(err, w){

					//send email confirmation
					go rabbitmq.BroadcastNewSubEvent(*sm)
					w.WriteHeader(200)
				}
			}
		}
	}

	w.WriteHeader(400)
}

func (s *Subscription) asyncRequest(body *models.Subscriber, c *stripe.Customer, t *stripe.Token, e error){

	s.wg.Add(1)
	go func(wg *sync.WaitGroup){
		defer wg.Done()
		c, e = s.CustomerClient.Put(body.Customer)
	}(s.wg)

	s.wg.Add(1)
	go func(wg *sync.WaitGroup){
		defer wg.Done()
		t, e = s.TokenClient.Put(body.PaymentDetails)
	}(s.wg)
}