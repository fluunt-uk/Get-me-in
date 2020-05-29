package event_driven

import (
	"github.com/ProjectReferral/Get-me-in/customer-api/configs"
	"github.com/ProjectReferral/Get-me-in/customer-api/internal/service"
	t "github.com/ProjectReferral/Get-me-in/customer-api/lib/hermes/templates"
	"github.com/ProjectReferral/Get-me-in/customer-api/models"
	"github.com/ProjectReferral/Get-me-in/queueing-api/client"
	queue_models "github.com/ProjectReferral/Get-me-in/queueing-api/client/models"
	"log"
	"net/http"
)

type MsgHandler struct {
	emailService *service.EmailService
}

func (r *MsgHandler) InjectService(s *service.EmailService) {
	r.emailService = s
}
//how we want to handle the incoming message
func (r *MsgHandler) HandleRabbitMessage(qm *queue_models.QueueMessage, err error, qc client.QueueClient) (queue_models.SubscribeMessage, client.HttpReponse) {

	if err != nil {
		log.Printf("error converting message [%s]", err)
	} else {
		log.Printf("message is: %+v", *qm)
	}
	var sm queue_models.SubscribeMessage

	p := models.IncomingData{}
	t.ToStruct(qm.Body, &p)

	r.emailService.SendEmail(qm.Body)

	if false {
		sm = queue_models.MessageReject{
			SubID:  	 *Store.GetSubscriber(configs.VERIFY_EMAIL_Q),
			ID:      	qm.ID,
			Requeue: 	true,
		}
	} else {
		sm = queue_models.MessageAcknowledge{
			SubID:       *Store.GetSubscriber(configs.VERIFY_EMAIL_Q),
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
