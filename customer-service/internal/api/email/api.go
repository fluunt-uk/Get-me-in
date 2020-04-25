package email

import (
	"net/http"
	"github.com/ProjectReferral/Get-me-in/customer-service/internal/event-driven"
)


func SendEmail(w http.ResponseWriter, r *http.Request) {
	event_driven.EmailBuilder.CreateEmail(w, r)
}
