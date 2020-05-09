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
	"time"
)

type Subscription struct {
	CustomerClient customer.Builder
	SubClient sub.Builder
	TokenClient token.Builder
	CardClient card.Builder
}

type asyncResponse struct {
	c *stripe.Customer
	t *stripe.Token
	err error
}

func (s *Subscription) SubscribeToPremiumPlan(w http.ResponseWriter, r *http.Request){

	wg := &sync.WaitGroup{}

	body := &models.Subscriber{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if !stripe_api.HandleError(err, w) {

		a := &asyncResponse{}
		//create new customer and token
		s.asyncRequest(body, a, wg)

		if !stripe_api.HandleError(err, w) {

			time.Sleep(configs.THROTTLE)
			//link the card with customer
			_, cardErr := s.CardClient.Put(a.c, a.t)

			if !stripe_api.HandleError(cardErr, w) {

				time.Sleep(configs.THROTTLE)
				//create the sub and make payment
				sm, subErr := s.SubClient.Put(a.c, configs.PREMIUM_PLAN)

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

func (s *Subscription) asyncRequest(body *models.Subscriber,
	a *asyncResponse, wg *sync.WaitGroup){

	wg.Add(1)
	go func(){
		defer wg.Done()
		a.c, a.err = s.CustomerClient.Put(body.Customer)
	}()

	wg.Add(1)
	go func(){
		defer wg.Done()
		time.Sleep(configs.THROTTLE)
		a.t, a.err = s.TokenClient.Put(body.PaymentDetails)
	}()

	wg.Wait()
}