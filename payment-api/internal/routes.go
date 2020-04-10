package internal

import (
	"github.com/ProjectReferral/Get-me-in/account-api/configs"
	"github.com/ProjectReferral/Get-me-in/pkg/security"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func SetupEndpoints() {

	_router := mux.NewRouter()

	_router.HandleFunc("/createCustomer", createCustomer).Methods("POST")
	_router.HandleFunc("/getCustomer", wrapHandlerWithSpecialAuth(retrieveCustomer, configs.AUTH_AUTHENTICATED)).Methods("GET")
	_router.HandleFunc("/updateCustomer", wrapHandlerWithSpecialAuth(updateCustomer, configs.AUTH_AUTHENTICATED)).Methods("POST")
	_router.HandleFunc("/delCustomer", wrapHandlerWithSpecialAuth(deleteCustomer, configs.AUTH_AUTHENTICATED)).Methods("DELETE")
	_router.HandleFunc("/getAllCustomers", wrapHandlerWithSpecialAuth(listAllCustomers, configs.AUTH_AUTHENTICATED)).Methods("GET")

	log.Fatal(http.ListenAndServe(configs.PORT, _router))
}
