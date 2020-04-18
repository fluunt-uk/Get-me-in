package api

import (
	"github.com/gorilla/mux"
	"github.com/ProjectReferral/Get-me-in/marketing-api/configs"
	"github.com/ProjectReferral/Get-me-in/util"
	"log"
	"net/http"
)

func SetupEndpoints() {
	_router := mux.NewRouter()

	_router.HandleFunc("/test", TestFunc)

	_router.HandleFunc("/advert", util.WrapHandlerWithSpecialAuth(CreateAdvert, "")).Methods("PUT")
	_router.HandleFunc("/advert", util.WrapHandlerWithSpecialAuth(DeleteAdvert, "")).Methods("DELETE")
	_router.HandleFunc("/advert", util.WrapHandlerWithSpecialAuth(UpdateAdvert, "")).Methods("PATCH")
	_router.HandleFunc("/advert", util.WrapHandlerWithSpecialAuth(GetAdvert, "")).Methods("GET")

	log.Fatal(http.ListenAndServe(configs.PORT, _router))
}
