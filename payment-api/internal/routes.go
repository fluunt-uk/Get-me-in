package internal

import (
	"github.com/ProjectReferral/Get-me-in/payment-api/configs"
	"github.com/ProjectReferral/Get-me-in/payment-api/internal/api"
	"github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api/resources/customer"
	"github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api/resources/subscription"
	"github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api/resources/token"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func SetupEndpoints() {

	_router := mux.NewRouter()

	_router.HandleFunc("/customer/create", customer.CreateCustomer).Methods("POST")
	_router.HandleFunc("/customer/get", customer.RetrieveCustomer).Methods("GET")
	_router.HandleFunc("/customer/update", customer.UpdateCustomer).Methods("POST")
	_router.HandleFunc("/customer/delete", customer.DeleteCustomer).Methods("DELETE")
	_router.HandleFunc("/customer/getAll", customer.ListAllCustomers).Methods("GET")

	_router.HandleFunc("/sub/create", subscription.NewSub).Methods("POST")
	_router.HandleFunc("/sub/get", subscription.RetrieveSub).Methods("GET")
	_router.HandleFunc("/sub/update", subscription.UpdateSub).Methods("POST")
	_router.HandleFunc("/sub/cancel", subscription.CancelSub).Methods("DELETE")
	_router.HandleFunc("/sub/getAll", subscription.ListSubs).Methods("GET")

	_router.HandleFunc("/token/create", token.CreateToken).Methods("POST")
	_router.HandleFunc("/token/get", token.GetToken).Methods("GET")

	_router.HandleFunc("/card/create", api.CreateCard).Methods("POST")
	_router.HandleFunc("/card/get", api.GetCard).Methods("GET")
	_router.HandleFunc("/card/update", api.UpdateCard).Methods("POST")
	_router.HandleFunc("/card/delete", api.DeleteCard).Methods("DELETE")
	_router.HandleFunc("/card/getAll", api.GetAllCards).Methods("GET")

	log.Fatal(http.ListenAndServe(configs.PORT, _router))
}
