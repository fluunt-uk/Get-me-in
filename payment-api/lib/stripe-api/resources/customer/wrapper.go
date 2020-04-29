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
	CreateCustomer(http.ResponseWriter, *http.Request)
	GetCustomer(http.ResponseWriter, *http.Request)
	DeleteCustomer(http.ResponseWriter, *http.Request)
	UpdateCustomer(http.ResponseWriter, *http.Request)
	ListAllCustomers(http.ResponseWriter, *http.Request)
}

type Wrapper struct{}

func (cw *Wrapper) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	body := models.Customer{}
	err := json.NewDecoder(r.Body).Decode(&body)
	stripe_api.HandleError(err,w)
	params := &stripe.CustomerParams{
		Name: &body.Name,
		Email: &body.Email,
	}
	c, _ := customer.New(params)

	stripe_api.ReturnSuccessJSON(w, &c)
}

func (cw *Wrapper) GetCustomer(w http.ResponseWriter, r *http.Request) {
	body := models.Customer{}
	err := json.NewDecoder(r.Body).Decode(&body)
	stripe_api.HandleError(err,w)

	c, _ := customer.Get(body.Id, nil)

	stripe_api.ReturnSuccessJSON(w, &c)
}

func (cw *Wrapper) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
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

func (cw *Wrapper) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	body := models.Customer{}
	err := json.NewDecoder(r.Body).Decode(&body)
	stripe_api.HandleError(err,w)

	c, _ := customer.Del(body.Id, nil)

	stripe_api.ReturnSuccessJSON(w, &c)
}

func (cw *Wrapper) ListAllCustomers(w http.ResponseWriter, r *http.Request) {
	params := &stripe.CustomerListParams{}
	params.Filters.AddFilter("limit", "", "3")
	i := customer.List(params)
	for i.Next() {
		c := i.Customer()

		stripe_api.ReturnSuccessJSON(w, &c)
	}
}
