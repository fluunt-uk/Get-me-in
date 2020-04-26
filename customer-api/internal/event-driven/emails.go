package event_driven

import (
	"github.com/ProjectReferral/Get-me-in/customer-api/internal/smtp"
	"log"
	"net/http"
)

type EmailStruct struct {
}

type EmailBuilder interface{
	 CreateActionEmail(http.ResponseWriter, *http.Request)
	 CreateNotificationEmail(http.ResponseWriter, *http.Request)
	 CreateSubscriptionEmail(http.ResponseWriter, *http.Request)
}

var Emails EmailBuilder

func (c *EmailStruct) CreateActionEmail(w http.ResponseWriter, r *http.Request) {

	CheckNoBody(w, r)

	body := r.Body

	//template, body
	template, subject := smtp.BaseTypeActionEmail()

	smtp.SendEmail([]string{email}, subject, template)
	log.Printf("Email sent")

}


func (c *EmailStruct) CreateNotificationEmail(w http.ResponseWriter, r *http.Request) {

	CheckNoBody(w, r)

	body := r.Body
	template, subject := smtp.BaseTypeNotificationEmail()

	smtp.SendEmail([]string{email}, subject, template)
	log.Printf("Email sent")

}

func (c *EmailStruct) CreateSubscriptionEmail(w http.ResponseWriter, r *http.Request) {

	CheckNoBody(w, r)

	body := r.Body
	template, subject := smtp.BaseTypeSubscriptionEmail()

	smtp.SendEmail([]string{email}, subject, template)
	log.Printf("Email sent")

}

func CheckNoBody(w http.ResponseWriter, r *http.Request) {
	if r.ContentLength < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No body error"))
		return
	}
}
