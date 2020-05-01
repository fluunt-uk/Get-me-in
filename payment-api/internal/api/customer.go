package api

import (
	customer "github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api/resources/customer"
	"net/http"
)

var CustomerClient customer.Builder

func CreateCustomer(w http.ResponseWriter, r *http.Request)  {
	//CustomerClient.Put(w, r)
}

func GetCustomer(w http.ResponseWriter, r *http.Request)  {
	CustomerClient.Get(w, r)
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request)  {
	CustomerClient.Patch(w, r)
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request)  {
	CustomerClient.Del(w, r)
}

func GetAllCustomers(w http.ResponseWriter, r *http.Request)  {
	CustomerClient.GetBatch(w, r)
}