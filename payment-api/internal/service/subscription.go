package service

import (
	"encoding/json"
	"github.com/ProjectReferral/Get-me-in/payment-api/internal/api"
	"github.com/ProjectReferral/Get-me-in/payment-api/internal/models"
	stripe_api "github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api"
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
	stripe_api.HandleError(err,w)

	//create new customer and token
	AsyncRequest(&body, w)
	wg.Wait()

	if e != nil {
		w.WriteHeader(400)
		w.Write([]byte(e.Error()))
		return
	}

	//link the card with customer
	api.CardClient.Put(c ,t)

	//create the sub and make payment
	//api.SubClient.Put(c, "plan_type")
}

func AsyncRequest(body *models.Subscriber, w http.ResponseWriter){

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