package email

import (
	event_driven "github.com/ProjectReferral/Get-me-in/customer-api/internal/event-driven"
	"github.com/ProjectReferral/Get-me-in/queueing-api/client/models"
	"net/http"
)

func CustomSubscribe(w http.ResponseWriter, r *http.Request) {

	//TODO: get the data from the body

	event_driven.SubscribeTo(models.QueueSubscribe{})
}
