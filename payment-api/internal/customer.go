package internal

import (
	"encoding/json"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
	"net/http"
)

func createCustomer(w http.ResponseWriter, r *http.Request) {
	params := &stripe.CustomerParams{
		Description: stripe.String("My First Test Customer (created for API docs)"),
	}
	c, _ := customer.New(params)
	toString, err := json.Marshal(c)

	if !HandleError(err, w) {
		w.Write(toString)
		w.WriteHeader(http.StatusOK)
	}

}

func retrieveCustomer(w http.ResponseWriter, r *http.Request) {
	c, _ := customer.Get("cus_H4HfdRtWmkH717", nil)
	toString, err := json.Marshal(c)

	if !HandleError(err, w) {
		w.Write(toString)
		w.WriteHeader(http.StatusOK)
	}
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	params := &stripe.CustomerParams{}
	params.AddMetadata("order_id", "6735")
	c, _ := customer.Update(
		"cus_H4HfdRtWmkH717",
		params,
	)

	toString, err := json.Marshal(c)
	if !HandleError(err, w) {
		w.Write(toString)
		w.WriteHeader(http.StatusOK)
	}
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	c, _ := customer.Del("cus_H4HfdRtWmkH717", nil)
	toString, err := json.Marshal(c)
	if !HandleError(err, w) {
		w.Write(toString)
		w.WriteHeader(http.StatusOK)
	}
}

func listAllCustomers(w http.ResponseWriter, r *http.Request) {
	params := &stripe.CustomerListParams{}
	params.Filters.AddFilter("limit", "", "3")
	i := customer.List(params)
	for i.Next() {
		c := i.Customer()
		toString, err := json.Marshal(c)
		if !HandleError(err, w) {
			w.Write(toString)
			w.WriteHeader(http.StatusOK)
		}
	}
}
