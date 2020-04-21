package internal

import (
	"github.com/ProjectReferral/Get-me-in/auth-api/configs"
	"github.com/ProjectReferral/Get-me-in/auth-api/internal/api/auth"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func SetupEndpoints() {
	_router := mux.NewRouter()

	_router.HandleFunc("/auth", auth.VerifyCredentials).Methods("GET")
	_router.HandleFunc("/auth/temp", auth.IssueRegistrationTempToken).Methods("GET")
	//test response that can be used for testing the internal/responses
	_router.HandleFunc("/mock", auth.MockResponse).Methods("GET")

	log.Fatal(http.ListenAndServe(configs.PORT, _router))
}
