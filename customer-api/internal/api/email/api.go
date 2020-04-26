package email

import (
	"github.com/ProjectReferral/Get-me-in/customer-api/internal/event-driven"
	"net/http"
)

var Service event_driven.EmailBuilder

func SendActionEmail(w http.ResponseWriter, r *http.Request) {
	Service.CreateActionEmail(w, r)
}

func SendNotificationEmail(w http.ResponseWriter, r *http.Request) {
	Service.CreateNotificationEmail(w, r)
}

func SendSubscriptionEmail(w http.ResponseWriter, r *http.Request) {
	Service.CreateSubscriptionEmail(w, r)
}