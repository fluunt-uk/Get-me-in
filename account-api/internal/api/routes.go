package api

import (
	"github.com/ProjectReferral/Get-me-in/account-api/configs"
	"github.com/ProjectReferral/Get-me-in/account-api/internal/api/account"
	"github.com/ProjectReferral/Get-me-in/util"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func SetupEndpoints() {
	
	_router := mux.NewRouter()

	_router.HandleFunc("/test", account.TestFunc)

	//token with correct register claim allowed
	_router.HandleFunc("/account", util.WrapHandlerWithSpecialAuth(account.CreateUser, configs.AUTH_REGISTER)).Methods("PUT")

	////token with correct authenticated claim allowed
	//_router.HandleFunc("/account", util.WrapHandlerWithSpecialAuth(account.UpdateUser, configs.AUTH_AUTHENTICATED)).Methods("PATCH")
	//_router.HandleFunc("/account", util.WrapHandlerWithSpecialAuth(account.GetUser, configs.AUTH_AUTHENTICATED)).Methods("GET")
	//
	////no one should have access apart from super users
	//_router.HandleFunc("/account", util.WrapHandlerWithSpecialAuth(account.DeleteUser, configs.NO_ACCESS)).Methods("DELETE")
	//
	//_router.HandleFunc("/account/premium", util.WrapHandlerWithSpecialAuth(account.IsUserPremium, configs.AUTH_AUTHENTICATED)).Methods("GET")
	//
	////token with correct sign in claim allowed
	//_router.HandleFunc("/account/signin", util.WrapHandlerWithSpecialAuth(sign_in.Login, configs.AUTH_LOGIN)).Methods("POST")
	//
	////token verification happening under the function
	//_router.HandleFunc("/account/verify", account.VerifyEmail).Methods("POST")
	//_router.HandleFunc("/account/verify/resend", account.ResendVerification).Methods("POST")
	//
	//_router.HandleFunc("/account/advert", account.ResendVerification).Methods("GET")

	log.Fatal(http.ListenAndServe(configs.PORT, _router))
}
