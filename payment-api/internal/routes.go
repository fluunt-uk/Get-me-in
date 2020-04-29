package internal

import (
	"github.com/ProjectReferral/Get-me-in/payment-api/configs"
	"github.com/ProjectReferral/Get-me-in/payment-api/internal/api"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func SetupEndpoints() {

	_router := mux.NewRouter()

	_router.HandleFunc("/customer/create", api.CreateCustomer).Methods("POST")
	_router.HandleFunc("/customer/get", api.GetCustomer).Methods("GET")
	_router.HandleFunc("/customer/update", api.UpdateCustomer).Methods("POST")
	_router.HandleFunc("/customer/delete", api.DeleteCustomer).Methods("DELETE")
	_router.HandleFunc("/customer/getAll", api.GetAllCustomers).Methods("GET")

	_router.HandleFunc("/sub/create", api.CreateSub).Methods("POST")
	_router.HandleFunc("/sub/get", api.GetSub).Methods("GET")
	_router.HandleFunc("/sub/update", api.UpdateSub).Methods("POST")
	_router.HandleFunc("/sub/cancel", api.CancelSub).Methods("DELETE")
	_router.HandleFunc("/sub/getAll", api.GetAllSubs).Methods("GET")

	_router.HandleFunc("/token/create", api.CreateToken).Methods("POST")
	_router.HandleFunc("/token/get", api.GetToken).Methods("GET")

	_router.HandleFunc("/card/create", api.CreateCard).Methods("POST")
	_router.HandleFunc("/card/get", api.GetCard).Methods("GET")
	_router.HandleFunc("/card/update", api.UpdateCard).Methods("POST")
	_router.HandleFunc("/card/delete", api.DeleteCard).Methods("DELETE")
	_router.HandleFunc("/card/getAll", api.GetAllCards).Methods("GET")

	log.Fatal(http.ListenAndServe(configs.PORT, _router))
}
