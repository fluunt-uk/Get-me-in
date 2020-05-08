package card

import (
	"encoding/json"
	stripe_api "github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api"
	"github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api/resources/models"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
	"net/http"
)

type Builder interface {
	Put(models.Customer) (*stripe.Customer, error)
	Get(http.ResponseWriter, *http.Request)
	Del(http.ResponseWriter, *http.Request)
	Patch(http.ResponseWriter, *http.Request)
	GetBatch(http.ResponseWriter, *http.Request)
}

type Wrapper struct{}

func (cw *Wrapper) Put(m models.Customer) (*stripe.Customer, error){

	params := &stripe.CustomerParams{
		Name: &m.Name,
		Email: &m.Email,
	}
	c, err := customer.New(params)

	if err != nil {
		return nil, err
	}

	return c, nil
}

func (cw *Wrapper) Get(w http.ResponseWriter, r *http.Request) {
	body := models.Customer{}
	err := json.NewDecoder(r.Body).Decode(&body)
	stripe_api.HandleError(err,w)

	c, _ := customer.Get(body.Id, nil)

	stripe_api.ReturnSuccessJSON(w, &c)
}

func (cw *Wrapper) Patch(w http.ResponseWriter, r *http.Request) {
	body := models.Customer{}
	err := json.NewDecoder(r.Body).Decode(&body)
	stripe_api.HandleError(err,w)

	params := &stripe.CustomerParams{}
	params.AddMetadata("order_id", "00022")
	c, _ := customer.Update(
		body.Id,
		params,
	)

	stripe_api.ReturnSuccessJSON(w, &c)
}

func (cw *Wrapper) Del(w http.ResponseWriter, r *http.Request) {
	body := models.Customer{}
	err := json.NewDecoder(r.Body).Decode(&body)
	stripe_api.HandleError(err,w)

	c, _ := customer.Del(body.Id, nil)

	stripe_api.ReturnSuccessJSON(w, &c)
}

func (cw *Wrapper) GetBatch(w http.ResponseWriter, r *http.Request) {
	params := &stripe.CustomerListParams{}
	params.Filters.AddFilter("limit", "", "3")
	i := customer.List(params)
	for i.Next() {
		c := i.Customer()

		stripe_api.ReturnSuccessJSON(w, &c)
	}
}
