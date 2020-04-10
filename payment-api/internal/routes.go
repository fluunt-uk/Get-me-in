package internal

import (
	"github.com/ProjectReferral/Get-me-in/account-api/configs"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func SetupEndpoints() {

	_router := mux.NewRouter()

	_router.HandleFunc("/customer/create", createCustomer).Methods("POST")
	_router.HandleFunc("/customer/get", retrieveCustomer).Methods("GET")
	_router.HandleFunc("/customer/update", updateCustomer).Methods("POST")
	_router.HandleFunc("/customer/delete", deleteCustomer).Methods("DELETE")
	_router.HandleFunc("/customer/getAll", listAllCustomers).Methods("GET")

	log.Fatal(http.ListenAndServe(configs.PORT, _router))
}
