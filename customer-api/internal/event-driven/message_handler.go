package event_driven

import (
	"github.com/ProjectReferral/Get-me-in/customer-api/lib/hermes"
	"github.com/ProjectReferral/Get-me-in/queueing-api/client"
	queue_models "github.com/ProjectReferral/Get-me-in/queueing-api/client/models"
	"log"
	"net/http"
)

type MsgHandler struct {
	emailService 	hermes.EmailBuilder
}

func (r *MsgHandler) InjectService(builder hermes.EmailBuilder) {
	r.emailService = builder
}
//how we want to handle the incoming message
func (r *MsgHandler) HandleRabbitMessage(qm *queue_models.QueueMessage, err error, qc client.QueueClient) (queue_models.SubscribeMessage, client.HttpReponse) {

	//TODO: testing purposes only, we need to store all sub id's in memory
	var subID = queue_models.QueueSubscribeId{ID: "test-sub"}

	if err != nil {
		log.Printf("error converting message [%s]", err)
	} else {
		log.Printf("message is: %+v", *qm)
	}
	var sm queue_models.SubscribeMessage

	//do some processing with the qm.body
	//r.emailService.CreateActionEmail()

	//hermes.CheckBodyStatus()

	if false {
		sm = queue_models.MessageReject{
			SubID:   subID,
			ID:      qm.ID,
			Requeue: true,
		}
	} else {
		sm = queue_models.MessageAcknowledge{
			SubID:       subID,
			ID:          qm.ID,
			Acknowledge: true,
			Requeue:     false,
			Multiple:    false,
		}
	}

	return sm, response
}

//getting back from /acknowledge
func response(resp *http.Response, err error) {
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
			log.Printf("failed to post message %d %v", resp.StatusCode, resp)
		}
	}
}