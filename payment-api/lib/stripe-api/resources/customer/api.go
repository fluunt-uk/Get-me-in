package customer

import (
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
	"net/http"
)

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	params := &stripe.CustomerParams{
		Name: stripe.String("Hamza"),
		Description: stripe.String("My First Test Customer (created for API docs)"),
	}
	c, _ := customer.New(params)

	ReturnSuccessJSON(w, c)
}

func RetrieveCustomer(w http.ResponseWriter, r *http.Request) {
	c, _ := customer.Get("cus_H6Sh6OUR88nUKr", nil)

	ReturnSuccessJSON(w, c)
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	params := &stripe.CustomerParams{}
	params.AddMetadata("order_id", "00022")
	c, _ := customer.Update(
		"cus_H6Sh6OUR88nUKr",
		params,
	)

	ReturnSuccessJSON(w, c)
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	c, _ := customer.Del("cus_H6Sh6OUR88nUKr", nil)

	ReturnSuccessJSON(w, c)
}

func ListAllCustomers(w http.ResponseWriter, r *http.Request) {
	params := &stripe.CustomerListParams{}
	params.Filters.AddFilter("limit", "", "3")
	i := customer.List(params)
	for i.Next() {
		c := i.Customer()

		ReturnSuccessJSON(w, c)
	}
}
