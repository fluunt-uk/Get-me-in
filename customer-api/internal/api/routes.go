package api

import (
	"github.com/ProjectReferral/Get-me-in/customer-api/configs"
	"github.com/ProjectReferral/Get-me-in/customer-api/internal/api/email"
	"github.com/ProjectReferral/Get-me-in/customer-api/lib/hermes"
	"github.com/ProjectReferral/Get-me-in/queueing-api/client"
	"github.com/ProjectReferral/Get-me-in/queueing-api/client/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type EndpointBuilder struct {
	router       *mux.Router
	emailService hermes.EmailBuilder
}

func (r *EndpointBuilder) InjectService(builder hermes.EmailBuilder) {
	r.emailService = builder
}

func (r *EndpointBuilder) SetupRouter(route *mux.Router) {
	r.router = route
}

func (r *EndpointBuilder) SetupEndpoints() {

	r.router.HandleFunc("/subscribe", email.CustomSubscribe).Methods("POST")
}

func (r *EndpointBuilder) SetupSubscriptionEndpoint(dqc *client.DefaultQueueClient) {

	hc := &http.Client{
		Timeout: 5 * time.Second,
	}

	r.SetupEmailActionRoute(dqc, hc)

	log.Fatal(http.ListenAndServe(configs.PORT, r.router))
}

func (r *EndpointBuilder) SetupEmailActionRoute(dqc *client.DefaultQueueClient, hc *http.Client) {
	dqc.SetupRoute(r.router, "/email/action", hc, r.handle)
}

//how we want to handle the incoming message
func (r *EndpointBuilder) handle(qm *models.QueueMessage, err error, qc client.QueueClient) (models.SubscribeMessage, client.HttpReponse) {

	//TODO: testing purposes only, we need to store all sub id's in memory
	var subID = models.QueueSubscribeId{ID: "test-sub"}

	if err != nil {
		log.Printf("error converting message [%s]", err)
	} else {
		log.Printf("message is: %+v", *qm)
	}
	var sm models.SubscribeMessage

	//do some processing with the qm.body
	r.emailService.CreateActionEmail()

	//hermes.CheckBodyStatus()

	if false {
		sm = models.MessageReject{
			SubID:   subID,
			ID:      qm.ID,
			Requeue: true,
		}
	} else {
		sm = models.MessageAcknowledge{
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
