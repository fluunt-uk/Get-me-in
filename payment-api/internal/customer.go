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
	c, err := customer.New(params)
	toString, _ := json.Marshal(c)

	if err {

	}
	w.Write(toString)
	w.WriteHeader(http.StatusOK)
}

func retrieveCustomer() *stripe.Customer {
	c, _ := customer.Get("cus_H4HfdRtWmkH717", nil)

	return c
}

func updateCustomer() *stripe.Customer {
	params := &stripe.CustomerParams{}
	params.AddMetadata("order_id", "6735")
	c, _ := customer.Update(
		"cus_H4HfdRtWmkH717",
		params,
	)

	return c
}

func deleteCustomer() *stripe.Customer {
	c, _ := customer.Del("cus_H4HfdRtWmkH717", nil)

	return c
}

func listAllCustomers() *stripe.Customer {
	params := &stripe.CustomerListParams{}
	params.Filters.AddFilter("limit", "", "3")
	i := customer.List(params)
	for i.Next() {
		c := i.Customer()
		return c
	}

	return nil
}
