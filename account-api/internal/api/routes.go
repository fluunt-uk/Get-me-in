package api

import (
	"github.com/ProjectReferral/Get-me-in/account-api/configs"
	"github.com/gorilla/mux"
	"github.com/ProjectReferral/Get-me-in/util"
	"log"
	"net/http"
)

func SetupEndpoints() {

	_router := mux.NewRouter()

	_router.HandleFunc("/test", TestFunc)

	//token with correct register claim allowed
	_router.HandleFunc("/account", util.WrapHandlerWithSpecialAuth(CreateUser, configs.AUTH_REGISTER)).Methods("PUT")

	//token with correct authenticated claim allowed
	_router.HandleFunc("/account", util.WrapHandlerWithSpecialAuth(UpdateUser, configs.AUTH_AUTHENTICATED)).Methods("PATCH")
	_router.HandleFunc("/account", util.WrapHandlerWithSpecialAuth(GetUser, configs.AUTH_AUTHENTICATED)).Methods("GET")

	//token with correct sign in claim allowed
	_router.HandleFunc("/account/signin", util.WrapHandlerWithSpecialAuth(Login, configs.AUTH_LOGIN)).Methods("POST")

	//no one should have access apart from super users
	_router.HandleFunc("/account", util.WrapHandlerWithSpecialAuth(DeleteUser, configs.NO_ACCESS)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(configs.PORT, _router))
}
