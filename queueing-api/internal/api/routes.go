package api

import (
	"github.com/ProjectReferral/Get-me-in/queueing-api/configs"
	"github.com/gorilla/mux"
	"net/http"
	"log"
)

func wrapHandlerWithBodyCheck(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.ContentLength < 1 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("No body error!"))
		}else{
			handler(w,r)
		}
	}
}

func SetupEndpoints() {

	_router := mux.NewRouter()

	log.Printf("url %s\n",configs.BrokerUrl)
	_router.HandleFunc("/test", TestFunc).Methods("GET")

	//create queue
	_router.HandleFunc("/queue", wrapHandlerWithBodyCheck(CreateQueue)).Methods("POST")

	//create echange
	_router.HandleFunc("/exchange", wrapHandlerWithBodyCheck(CreateExchange)).Methods("POST")

	//put bind
	_router.HandleFunc("/bind", wrapHandlerWithBodyCheck(BindExchange)).Methods("PUT")

	//publish message
	_router.HandleFunc("/publish", wrapHandlerWithBodyCheck(PublishToExchange)).Methods("POST")

	//consume message
	_router.HandleFunc("/consume", wrapHandlerWithBodyCheck(ConsumeQueue)).Methods("POST")

	//subscribe queue
	_router.HandleFunc("/subscribe", wrapHandlerWithBodyCheck(SuscribeQueue)).Methods("POST")

	//unsubscribe queue
	_router.HandleFunc("/unsubscribe", wrapHandlerWithBodyCheck(UnSuscribeQueue)).Methods("POST")

	//acknowledge message
	_router.HandleFunc("/acknowledge", wrapHandlerWithBodyCheck(MessageAck)).Methods("POST")
	
	//reject message
	_router.HandleFunc("/reject", wrapHandlerWithBodyCheck(MessageReject)).Methods("POST")

	//dump data - requires password
	_router.HandleFunc("/dump", wrapHandlerWithBodyCheck(DumpData)).Methods("POST")

	log.Println("Service started")
	log.Fatal(http.ListenAndServe(configs.PORT, _router))
}