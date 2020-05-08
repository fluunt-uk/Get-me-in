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

	//TODO: these might not be thread safe
	c *stripe.Customer
	t *stripe.Token
	e error
}

func (s *Subscription) Init(){
	s.wg = &sync.WaitGroup{}
}

func (s *Subscription) SubscribeToPremiumPlan(w http.ResponseWriter, r *http.Request){

	body := &models.Subscriber{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if !stripe_api.HandleError(err, w) {

		//create new customer and token
		s.asyncRequest(body)
		s.wg.Wait()

		if !stripe_api.HandleError(s.e, w) {

			//link the card with customer
			s.RLock()
			_, cardErr := s.CardClient.Put(s.c, s.t)
			s.RUnlock()

			if !stripe_api.HandleError(cardErr, w) {

				s.RLock()
				//create the sub and make payment
				sm, subErr := s.SubClient.Put(s.c, configs.PREMIUM_PLAN)
				s.RUnlock()

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

func (s *Subscription) asyncRequest(body *models.Subscriber){

	s.wg.Add(1)
	go func(wg *sync.WaitGroup){
		defer wg.Done()
		s.Lock()
		s.c, s.e = s.CustomerClient.Put(body.Customer)
		s.Unlock()
	}(s.wg)

	s.wg.Add(1)
	go func(wg *sync.WaitGroup){
		defer wg.Done()
		s.Lock()
		s.t, s.e = s.TokenClient.Put(body.PaymentDetails)
		s.Unlock()
	}(s.wg)
}