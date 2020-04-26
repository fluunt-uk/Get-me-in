package api

import (
	"github.com/ProjectReferral/Get-me-in/customer-api/configs"
	s "github.com/ProjectReferral/Get-me-in/customer-api/internal/api/email"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var Router = *mux.NewRouter()

func SetupEndpoints() {

	Router.HandleFunc("/email/notification", s.SendNotificationEmail).Methods("POST")
	Router.HandleFunc("/email/subscription", s.SendSubscriptionEmail).Methods("POST")

	log.Fatal(http.ListenAndServe(configs.PORT, &Router))
}
