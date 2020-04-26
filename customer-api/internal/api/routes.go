package api

import (
	"github.com/ProjectReferral/Get-me-in/customer-api/configs"
	s "github.com/ProjectReferral/Get-me-in/customer-api/internal/api/email"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func SetupEndpoints() {
	_router := mux.NewRouter()

	_router.HandleFunc("/email/action", s.SendActionEmail).Methods("POST")
	_router.HandleFunc("/email/notification", s.SendNotificationEmail).Methods("POST")
	_router.HandleFunc("/email/subscription", s.SendSubscriptionEmail).Methods("POST")
	log.Fatal(http.ListenAndServe(configs.PORT, _router))
}
