package event_driven

import (
	"github.com/ProjectReferral/Get-me-in/customer-api/lib/hermes"
	t "github.com/ProjectReferral/Get-me-in/customer-api/lib/hermes/templates"
	"github.com/ProjectReferral/Get-me-in/queueing-api/client"
	queue_models "github.com/ProjectReferral/Get-me-in/queueing-api/client/models"
	"log"
	"net/http"
	"strings"
)

type MsgHandler struct {
	emailService 	hermes.EmailBuilder
}

type TemplateType struct {
	template 		string
}

func (r *MsgHandler) InjectService(builder hermes.EmailBuilder) {
	r.emailService = builder
}
//how we want to handle the incoming message
func (r *MsgHandler) HandleRabbitMessage(qm *queue_models.QueueMessage, err error, qc client.QueueClient) (queue_models.SubscribeMessage, client.HttpReponse) {

	if err != nil {
		log.Printf("error converting message [%s]", err)
	} else {
		log.Printf("message is: %+v", *qm)
	}
	var sm queue_models.SubscribeMessage

	p := TemplateType{}
	t.ToStruct(qm.Body, &p)

	if strings.Contains(p.template, "action") {
		r.emailService.CreateActionEmail(qm.Body)
	} else if strings.Contains(p.template, "payment") {
		r.emailService.CreateSubscriptionEmail(qm.Body)
	} else {
		r.emailService.CreateNotificationEmail(qm.Body)
	}

	//hermes.CheckBodyStatus()

	if false {
		sm = queue_models.MessageReject{
			SubID:   subcriberStore["new-user-verify-email"],
			ID:      qm.ID,
			Requeue: true,
		}
	} else {
		sm = queue_models.MessageAcknowledge{
			SubID:       subcriberStore["new-user-verify-email"],
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
