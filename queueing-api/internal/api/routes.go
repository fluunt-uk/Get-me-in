package api

import (
	"github.com/ProjectReferral/Get-me-in/queueing-api/configs"
	"github.com/ProjectReferral/Get-me-in/queueing-api/client"
	"github.com/ProjectReferral/Get-me-in/queueing-api/client/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
	"encoding/json"
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

	//unsubscribe queue
	_router.HandleFunc("/dump", wrapHandlerWithBodyCheck(DumpData)).Methods("POST")

	dqc := &client.DefaultQueueClient{}
	dqc.SetupURL("http://localhost:5004")
	setupRoute(_router, dqc)
	go testSub(dqc)
	
	log.Println("Service started")
	log.Fatal(http.ListenAndServe(configs.PORT, _router))
}

func handle(qm *models.QueueMessage, err error, qc client.QueueClient){
	if err != nil {
		log.Printf("error converting message ",err)
	}else{
		log.Printf("message is: %+v",*qm)
	}
}

func setupRoute(router *mux.Router, dqc *client.DefaultQueueClient){
	dqc.SetupRoute(router,"/receive",handle)
}

func testSub(dqc *client.DefaultQueueClient){
	defer log.Printf("[end] testSub")
	time.Sleep(5 * time.Second)
	hc := &http.Client{
		Timeout: 5 * time.Second,
	}
	subscribe := models.QueueSubscribe {
		URL: "http://localhost:5004/receive",
		Name: "test-queue",
		Consumer: "",
		Exclusive: false,
		NoLocal: false,
		NoWait: false,
		MaxRetry: 3,
		Timeout: 5 * time.Second,
		Qos: models.QueueQos {
			PrefetchCount: 0,
			PrefetchSize: 0,
		},
	}
	resp,err := dqc.Subscribe(hc, subscribe)
	if err == nil || resp != nil{
		defer resp.Body.Close()
		subId := models.QueueSubscribeId{}
		jsonError := json.NewDecoder(resp.Body).Decode(&subId)
		if jsonError != nil{
			log.Printf("failed to make json read [%+v] status[%d]",jsonError,resp.StatusCode)
		}else{
			log.Printf("body: %+v",subId)
			
			go func(){
				time.Sleep(5 * time.Second)
				log.Println("testing....")
				testUnSub(hc,subId,dqc)
			}()
		}
	}else{
		log.Printf("error sending POST err[%+v] resp[%+v]",err, resp)
	}
}

func testUnSub(hc *http.Client, subID models.QueueSubscribeId, dqc *client.DefaultQueueClient){
	defer log.Printf("[end] testUnSub")
	resp,err := dqc.UnSubscribe(hc, subID)
	if err == nil || resp != nil{
		defer resp.Body.Close()
		log.Printf("unsunbsribe status: [%d]",resp.StatusCode)
	}else{
		log.Printf("error sending POST err[%+v] resp[%+v]",err, resp)
	}
}
