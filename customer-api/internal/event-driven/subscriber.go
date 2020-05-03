package event_driven

import (
	"encoding/json"
	"github.com/ProjectReferral/Get-me-in/queueing-api/client"
	"github.com/ProjectReferral/Get-me-in/queueing-api/client/models"
	"log"
	"net/http"
	"time"
)

var Client client.QueueClient

var SubID = models.QueueSubscribeId{}

var subcriberStore = make(map[string]models.QueueSubscribeId)

func SubscribeTo(sub models.QueueSubscribe){

	hc := &http.Client{Timeout: 5 * time.Second}

	defer log.Printf("[end] testSub")
	time.Sleep(5 * time.Second)

	resp,err := Client.Subscribe(hc, sub)
	if err == nil || resp != nil{
		defer resp.Body.Close()
		jsonError := json.NewDecoder(resp.Body).Decode(&SubID)
		if jsonError != nil{
			log.Printf("failed to make json read [%+v] status[%d]",jsonError,resp.StatusCode)
		}else{
			log.Printf("body: %+v",SubID)
			subcriberStore["new-user-verify-email"] = SubID
		}
	}else{
		log.Printf("error sending POST err[%+v] resp[%+v]",err, resp)
	}

}
