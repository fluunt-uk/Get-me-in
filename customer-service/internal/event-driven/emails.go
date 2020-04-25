package event_driven

import (
	"github.com/ProjectReferral/Get-me-in/customer-service/internal/smtp"
	"github.com/ProjectReferral/Get-me-in/customer-service/lib/hermes/templates"
	"log"
	"net/http"
)

//implement only the necessary methods for each repository
//available to be consumed by the API
type EmailBuilder interface{
	 CreateEmail(http.ResponseWriter, *http.Request)
}
//interface with the implemented methods will be injected in this variable
var Emails EmailBuilder

func CreateEmail(w http.ResponseWriter, r *http.Request) {

	if r.ContentLength < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No body error"))
		return
	}
	// Need to unmarshal body to struct and use types from there
	// Will need base_type & email_type
	body := r.Body

	// Will need correct stuff from body to pass into params
	// Basetype, template, body
	template, email, subject := templates.GenerateHTMLTemplate()

	smtp.SendEmail([]string{email}, subject, template)
	log.Printf("Email sent")

}

