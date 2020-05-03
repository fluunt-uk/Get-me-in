package api

import (
	"github.com/ProjectReferral/Get-me-in/customer-api/configs"
	"github.com/ProjectReferral/Get-me-in/customer-api/internal/api/email"
	event_driven "github.com/ProjectReferral/Get-me-in/customer-api/internal/event-driven"
	"github.com/ProjectReferral/Get-me-in/queueing-api/client"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type EndpointBuilder struct {
	router       	*mux.Router
	dqc				*client.DefaultQueueClient
	msg 			*event_driven.MsgHandler
}

func (r *EndpointBuilder) SetupRouter(route *mux.Router) {
	r.router = route
}

func (r *EndpointBuilder) SetupMsgHandler(mh *event_driven.MsgHandler) {
	r.msg = mh
}

func (r *EndpointBuilder) SetQueueClient(dqc *client.DefaultQueueClient) {
	r.dqc = dqc
}

func (r *EndpointBuilder) SetupEndpoints() {

	r.router.HandleFunc("/subscribe", email.CustomSubscribe).Methods("POST")
}

func (r *EndpointBuilder) SetupSubscriptionEndpoint() {

	hc := &http.Client{
		Timeout: 5 * time.Second,
	}

	r.SetupEmailActionRoute(hc)

	log.Fatal(http.ListenAndServe(configs.PORT, r.router))
}

func (r *EndpointBuilder) SetupEmailActionRoute(hc *http.Client) {
	r.dqc.SetupRoute(r.router, "/email/action", hc, r.msg.HandleRabbitMessage)
}

