package internal

import (
	"github.com/rs/cors"
	//"github.com/gorilla/handlers"
	"github.com/ProjectReferral/Get-me-in/auth-api/configs"
	"github.com/ProjectReferral/Get-me-in/auth-api/internal/api/auth"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func SetupEndpoints() {
	_router := mux.NewRouter()

	_router.HandleFunc("/auth", auth.VerifyCredentials)
	_router.HandleFunc("/auth/temp", auth.IssueRegistrationTempToken).Methods("POST")
	//test response that can be used for testing the internal/responses
	_router.HandleFunc("/mock", auth.MockResponse)

	c := cors.New(cors.Options{
    AllowedMethods: []string{"POST"},
    AllowedOrigins: []string{"*"},
    AllowCredentials: true,
    AllowedHeaders: []string{"Authorization", "Content-Type","Origin","Accept", "Accept-Encoding", "Accept-Language", "Host", "Connection", "Referer", "Sec-Fetch-Mode", "User-Agent", "Access-Control-Request-Headers", "Access-Control-Request-Method: "},
    OptionsPassthrough: true,
})



handler := c.Handler(_router)
log.Fatal(http.ListenAndServe(configs.PORT,handler))
}
