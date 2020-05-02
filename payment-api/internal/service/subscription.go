package service

import (
	"encoding/json"
	"fmt"
	"github.com/ProjectReferral/Get-me-in/payment-api/configs"
	"github.com/ProjectReferral/Get-me-in/payment-api/internal/api"
	"github.com/ProjectReferral/Get-me-in/payment-api/internal/models"
	stripe_api "github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api"
	"github.com/ProjectReferral/Get-me-in/pkg/http_lib"
	"github.com/stripe/stripe-go"
	"net/http"
	"sync"
)

var (
	wg sync.WaitGroup
	c *stripe.Customer
	t *stripe.Token
	e error
	)


func SubscribeToPremiumPlan(w http.ResponseWriter, r *http.Request){

	body := models.Subscriber{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if !stripe_api.HandleError(err, w) {

		//create new customer and token
		AsyncRequest(&body)
		wg.Wait()

		if !stripe_api.HandleError(e, w) {

			//link the card with customer
			_, cardErr := api.CardClient.Put(c, t)
			if !stripe_api.HandleError(cardErr, w) {

				//create the sub and make payment
				_, subErr := api.SubClient.Put(c, configs.PREMIUM_PLAN)

				b, _ := json.Marshal(models.ChangeRequest{
					Field:   "premium",
					NewBool: true,
					Type:    3,
				})

				res, err := http_lib.Patch(configs.ACCOUNT_API_PREMIUM, b,
					map[string]string{"Authorization": r.Header.Get("Authorization")})

				fmt.Println(res, err)

				if !stripe_api.HandleError(subErr, w){
					w.WriteHeader(200)
				}
			}
		}
	}
	//go rabbitmq.BroadcastNewSubEvent(s)

	w.WriteHeader(400)
}

func AsyncRequest(body *models.Subscriber){

	wg.Add(1)
	go func(wg *sync.WaitGroup){
		defer wg.Done()
		c, e = api.CustomerClient.Put(body.Customer)
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup){
		defer wg.Done()
		t, e = api.TokenClient.Put(body.PaymentDetails)
	}(&wg)
}