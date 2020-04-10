package internal

import (
	"github.com/ProjectReferral/Get-me-in/account-api/configs"
	"github.com/ProjectReferral/Get-me-in/payment-api/internal/stripe-api"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func SetupEndpoints() {

	_router := mux.NewRouter()

	_router.HandleFunc("/customer/create", stripe_api.CreateCustomer).Methods("POST")
	_router.HandleFunc("/customer/get", stripe_api.RetrieveCustomer).Methods("GET")
	_router.HandleFunc("/customer/update", stripe_api.UpdateCustomer).Methods("POST")
	_router.HandleFunc("/customer/delete", stripe_api.DeleteCustomer).Methods("DELETE")
	_router.HandleFunc("/customer/getAll", stripe_api.ListAllCustomers).Methods("GET")

	_router.HandleFunc("/sub/create", stripe_api.CreateSub).Methods("POST")
	_router.HandleFunc("/sub/get", stripe_api.RetrieveSub).Methods("GET")
	_router.HandleFunc("/sub/update", stripe_api.UpdateSub).Methods("POST")
	_router.HandleFunc("/sub/cancel", stripe_api.CancelSub).Methods("DELETE")
	_router.HandleFunc("/sub/getAll", stripe_api.ListSubs).Methods("GET")

	log.Fatal(http.ListenAndServe(configs.PORT, _router))
}
