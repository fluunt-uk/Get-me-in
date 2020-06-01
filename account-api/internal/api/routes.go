package api

import (
	"github.com/ProjectReferral/Get-me-in/account-api/configs"
	"github.com/ProjectReferral/Get-me-in/account-api/internal/api/account"
	account_advert "github.com/ProjectReferral/Get-me-in/account-api/internal/api/account-advert"
	sign_in "github.com/ProjectReferral/Get-me-in/account-api/internal/api/sign-in"
	"github.com/ProjectReferral/Get-me-in/pkg/security"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func SetupEndpoints() {
	
	_router := mux.NewRouter()

	_router.HandleFunc("/test", account.TestFunc)

	//token with correct register claim allowed
	_router.HandleFunc("/account", security.WrapHandlerWithSpecialAuth(account.CreateUser, configs.AUTH_REGISTER)).Methods("PUT")

	//token with correct authenticated claim allowed
	_router.HandleFunc("/account", security.WrapHandlerWithSpecialAuth(account.UpdateUser, configs.AUTH_AUTHENTICATED)).Methods("PATCH")
	_router.HandleFunc("/account", security.WrapHandlerWithSpecialAuth(account.GetUser, configs.AUTH_AUTHENTICATED)).Methods("GET")

	//no one should have access apart from super users
	//_router.HandleFunc("/account", util.WrapHandlerWithSpecialAuth(account.DeleteUser, configs.NO_ACCESS)).Methods("DELETE")

	_router.HandleFunc("/account/premium", security.WrapHandlerWithSpecialAuth(account.IsUserPremium, configs.AUTH_AUTHENTICATED)).Methods("GET")

	//token with correct sign in claim allowed
	_router.HandleFunc("/account/signin", security.WrapHandlerWithSpecialAuth(sign_in.Login, configs.AUTH_LOGIN)).Methods("POST")

	//token verification happening under the function
	_router.HandleFunc("/account/verify", security.WrapHandlerWithSpecialAuth(account.VerifyEmail, "")).Methods("POST")
	_router.HandleFunc("/account/verify/resend", security.WrapHandlerWithSpecialAuth(account.ResendVerification, "")).Methods("POST")

	//user must be authenticated before access this endpoint
	_router.HandleFunc("/account/advert", security.WrapHandlerWithSpecialAuth(account_advert.GetAllAdverts, configs.AUTH_AUTHENTICATED)).Methods("GET")

	_router.HandleFunc("/log", displayLog).Methods("GET")

	log.Fatal(http.ListenAndServe(configs.PORT, _router))
}

func displayLog(w http.ResponseWriter, r *http.Request){
	b, _ := ioutil.ReadFile("logs/accountAPI_log.txt")

	w.Write(b)
}