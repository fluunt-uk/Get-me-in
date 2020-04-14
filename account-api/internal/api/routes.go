package api

import (
	"github.com/ProjectReferral/Get-me-in/account-api/configs"
	"github.com/ProjectReferral/Get-me-in/account-api/internal/api/account"
	sign_in "github.com/ProjectReferral/Get-me-in/account-api/internal/api/sign-in"
	"github.com/ProjectReferral/Get-me-in/pkg/security"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//Parses the authentication token and validates against the @claim
//Some tokens can only authenticate with specific endpoints
func wrapHandlerWithSpecialAuth(handler http.HandlerFunc, claim string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		a := req.Header.Get("Authorization")

		//empty header
		if a == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("No Authorization JTW!!"))
			return
		}

		//not empty header and token is valid
		if a != "" && security.VerifyTokenWithClaim(a, claim) {
			handler(w, req)
			return
		}

		//not empty header and token is invalid
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func SetupEndpoints() {

	_router := mux.NewRouter()

	_router.HandleFunc("/test", account.TestFunc)

	//token with correct register claim allowed
	_router.HandleFunc("/account", wrapHandlerWithSpecialAuth(account.CreateUser, configs.AUTH_REGISTER)).Methods("PUT")

	//token with correct authenticated claim allowed
	_router.HandleFunc("/account", wrapHandlerWithSpecialAuth(account.UpdateUser, configs.AUTH_AUTHENTICATED)).Methods("PATCH")
	_router.HandleFunc("/account", wrapHandlerWithSpecialAuth(account.GetUser, configs.AUTH_AUTHENTICATED)).Methods("GET")

	//no one should have access apart from super users
	_router.HandleFunc("/account", wrapHandlerWithSpecialAuth(account.DeleteUser, configs.NO_ACCESS)).Methods("DELETE")

	//no one should have access apart from super users
	_router.HandleFunc("/account/premium", wrapHandlerWithSpecialAuth(account.IsUserPremium, configs.AUTH_AUTHENTICATED)).Methods("GET")

	//token with correct sign in claim allowed
	_router.HandleFunc("/account/signin", wrapHandlerWithSpecialAuth(sign_in.Login, configs.AUTH_LOGIN)).Methods("POST")

	//token verification happening under the function
	_router.HandleFunc("/account/verify", account.VerifyEmail).Methods("GET")

	log.Fatal(http.ListenAndServe(configs.PORT, _router))
}
