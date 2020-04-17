package customer

import (
	"github.com/ProjectReferral/Get-me-in/payment-api/configs"
	stripe_api "github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
	"net/http"
)

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	params := &stripe.CustomerParams{
		Name: stripe.String("Hamza 2"),
		Description: stripe.String("My First Test Customer (created for API docs)"),
	}
	c, _ := customer.New(params)

	configs.StripeObjects["customer"] = c
	stripe_api.ReturnSuccessJSON(w,"customer")
}

func RetrieveCustomer(w http.ResponseWriter, r *http.Request) {
	c, _ := customer.Get("cus_H6Sh6OUR88nUKr", nil)

	configs.StripeObjects["customer"] = c
	stripe_api.ReturnSuccessJSON(w,"customer")
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	params := &stripe.CustomerParams{}
	params.AddMetadata("order_id", "00022")
	c, _ := customer.Update(
		"cus_H6Sh6OUR88nUKr",
		params,
	)

	configs.StripeObjects["customer"] = c
	stripe_api.ReturnSuccessJSON(w,"customer")
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	c, _ := customer.Del("cus_H6Sh6OUR88nUKr", nil)

	configs.StripeObjects["customer"] = c
	stripe_api.ReturnSuccessJSON(w,"customer")
}

func ListAllCustomers(w http.ResponseWriter, r *http.Request) {
	params := &stripe.CustomerListParams{}
	params.Filters.AddFilter("limit", "", "3")
	i := customer.List(params)
	for i.Next() {
		c := i.Customer()

		configs.StripeObjects["customer"] = c
		stripe_api.ReturnSuccessJSON(w,"customer")
	}
}
