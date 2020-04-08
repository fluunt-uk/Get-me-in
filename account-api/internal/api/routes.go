package api

import (
	"github.com/ProjectReferral/Get-me-in/account-api/configs"
	"github.com/ProjectReferral/Get-me-in/pkg/security"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var email = ""

//Parses the authentication token and validates against the @claim
//Some tokens can only authenticate with specific endpoints
func wrapHandlerWithSpecialAuth(handler http.HandlerFunc, claim string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		a := req.Header.Get("Authorization")

		verified, claims := security.VerifyTokenWithClaim(a, claim)
		email = claims.Subject // get the email from the subject

		if a != "" && verified {
			handler(w, req)
			return
		}

		w.WriteHeader(http.StatusUnauthorized)
	}
}

// return the email value from the JWT
func getEmailValue() string {
	return email
}

func SetupEndpoints() {

	_router := mux.NewRouter()

	_router.HandleFunc("/test", TestFunc)

	_router.HandleFunc("/account", wrapHandlerWithSpecialAuth(CreateUser, configs.AUTH_REGISTER)).Methods("PUT")
	_router.HandleFunc("/account", wrapHandlerWithSpecialAuth(DeleteUser, configs.AUTH_AUTHENTICATED)).Methods("DELETE")
	_router.HandleFunc("/account", wrapHandlerWithSpecialAuth(UpdateUser, configs.AUTH_AUTHENTICATED)).Methods("PATCH")
	_router.HandleFunc("/account", wrapHandlerWithSpecialAuth(GetUser, configs.AUTH_AUTHENTICATED)).Methods("GET")
	//_router.HandleFunc("/account", GetUsers).Methods("GET")
	_router.HandleFunc("/account/verify", wrapHandlerWithSpecialAuth(Login, configs.AUTH_LOGIN)).Methods("POST")

	log.Fatal(http.ListenAndServe(configs.PORT, _router))
}
