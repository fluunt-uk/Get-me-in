package email

import (
	"net/http"
	"github.com/ProjectReferral/Get-me-in/customer-api/internal/event-driven"
)


func SendActionEmail(w http.ResponseWriter, r *http.Request) {
	event_driven.EmailBuilder.CreateActionEmail(w, r)
}

func SendNotificationEmail(w http.ResponseWriter, r *http.Request) {
	event_driven.EmailBuilder.CreateNotificationEmail(w, r)
}

func SendSubscriptionEmail(w http.ResponseWriter, r *http.Request) {
	event_driven.EmailBuilder.CreateSubscriptionEmail(w, r)
}