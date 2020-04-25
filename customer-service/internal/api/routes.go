package api

import (
	"github.com/ProjectReferral/Get-me-in/customer-service/configs"
	s "github.com/ProjectReferral/Get-me-in/customer-service/internal/api/email"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func SetupEndpoints() {
	_router := mux.NewRouter()

	_router.HandleFunc("/email", s.SendEmail).Methods("POST")
	log.Fatal(http.ListenAndServe(configs.PORT, _router))
}
