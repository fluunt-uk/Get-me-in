package internal

import (
	"github.com/ProjectReferral/Get-me-in/payment-api/configs"
	"github.com/ProjectReferral/Get-me-in/payment-api/internal/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var SS service.Subscription

func SetupEndpoints() {

	_router := mux.NewRouter()

	_router.HandleFunc("/premium/subscribe", SS.SubscribeToPremiumPlan).Methods("POST")

	log.Fatal(http.ListenAndServe(configs.PORT, _router))
}
