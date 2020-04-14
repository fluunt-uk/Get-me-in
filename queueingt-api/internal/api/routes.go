package api

import (
	"github.com/ProjectReferral/Get-me-in/queueingt-api/configs"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func SetupEndpoints() {

	configs.BrokerUrl = os.Getenv("BROKERURL")
	_router := mux.NewRouter()

	log.Printf("url %s\n",configs.BrokerUrl)
	_router.HandleFunc("/test", TestFunc)

	//create queue
	_router.HandleFunc("/queue", (CreateQueue)).Methods("POST")
	
	//create queue
	_router.HandleFunc("/exchange", (CreateExchange)).Methods("POST")
	
	//create queue
	_router.HandleFunc("/bind", (BindExchange)).Methods("PUT")
	
	//create queue
	_router.HandleFunc("/publish", (PublishToExchange)).Methods("POST")

	log.Fatal(http.ListenAndServe(configs.PORT, _router))
}
