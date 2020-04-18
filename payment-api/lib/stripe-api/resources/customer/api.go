package customer

import (
	"encoding/json"
	stripe_api "github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api"
	"github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api/resources/models"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
	"net/http"
)

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
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

func RetrieveCustomer(w http.ResponseWriter, r *http.Request) {
	body := models.Customer{}
	err := json.NewDecoder(r.Body).Decode(&body)
	stripe_api.HandleError(err,w)

	c, _ := customer.Get(body.Id, nil)

	stripe_api.ReturnSuccessJSON(w, &c)
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
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

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	body := models.Customer{}
	err := json.NewDecoder(r.Body).Decode(&body)
	stripe_api.HandleError(err,w)

	c, _ := customer.Del(body.Id, nil)

	stripe_api.ReturnSuccessJSON(w, &c)
}

func ListAllCustomers(w http.ResponseWriter, r *http.Request) {
	params := &stripe.CustomerListParams{}
	params.Filters.AddFilter("limit", "", "3")
	i := customer.List(params)
	for i.Next() {
		c := i.Customer()

		stripe_api.ReturnSuccessJSON(w, &c)
	}
}
