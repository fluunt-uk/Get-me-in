package event_driven

import (
	"encoding/json"
	"github.com/ProjectReferral/Get-me-in/customer-api/configs"
	"github.com/ProjectReferral/Get-me-in/queueing-api/client"
	"github.com/ProjectReferral/Get-me-in/queueing-api/client/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

var Client client.QueueClient

var b = false

var subID = models.QueueSubscribeId{}

func Subscribe(dqc *client.DefaultQueueClient, hc *http.Client){
	defer log.Printf("[end] testSub")
	time.Sleep(5 * time.Second)
	subscribe := models.QueueSubscribe {
		//endpoint that will be consuming the messages
		URL: configs.SUB_ACTION_EMAIL,
		Name: "test-queue",
		Consumer: "",
		Exclusive: false,
		NoLocal: false,
		NoWait: false,
		MaxRetry: 0,
		Timeout: 5 * time.Second,
		Qos: models.QueueQos {
			PrefetchCount: 0,
			PrefetchSize: 0,
		},
	}

	resp,err := dqc.Subscribe(hc, subscribe)
	if err == nil || resp != nil{
		defer resp.Body.Close()
		jsonError := json.NewDecoder(resp.Body).Decode(&subID)
		if jsonError != nil{
			log.Printf("failed to make json read [%+v] status[%d]",jsonError,resp.StatusCode)
		}else{
			log.Printf("body: %+v",subID)
			go func(){
				time.Sleep(10 * time.Second)
				//testUnSub(hc,subId,dqc)
			}()
		}
	}else{
		log.Printf("error sending POST err[%+v] resp[%+v]",err, resp)
	}

}

func SetupRoute(router *mux.Router, hc *http.Client, dqc *client.DefaultQueueClient){
	dqc.SetupRoute(router,"/email/action", hc, handle)
}

//how we want to handle the incoming message
func handle(qm *models.QueueMessage, err error, qc client.QueueClient) (models.SubscribeMessage, client.HttpReponse) {
	if err != nil {
		log.Printf("error converting message [%s]",err)
	}else{
		log.Printf("message is: %+v",*qm)
	}
	var sm models.SubscribeMessage

	//do some processing with the qm.body
	//email.Service.CreateActionEmail()

	if false {
		sm = models.MessageReject {
			SubID: subID,
			ID: qm.ID,
			Requeue: true,
		}
	}else{
		sm = models.MessageAcknowledge {
			SubID: subID,
			ID: qm.ID,
			Acknowledge: true,
			Requeue: false,
			Multiple: false,
		}
	}

	return sm, Response
}


//getting back from /acknowledge
func Response(resp *http.Response, err error){
	if err != nil {
		log.Println(err)
		return
	}
	if resp != nil {
		defer resp.Body.Close()
		switch resp.StatusCode {
		case 200:
			return
		case 403:
			log.Printf("Error: [403]")
		default:
			log.Printf("failed to post message %d %v",resp.StatusCode,resp)
		}
	}
}
