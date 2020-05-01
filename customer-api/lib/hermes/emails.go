package hermes

import (
	"github.com/ProjectReferral/Get-me-in/customer-api/internal/smtp"
	t "github.com/ProjectReferral/Get-me-in/customer-api/lib/hermes/templates"
	"github.com/ProjectReferral/Get-me-in/customer-api/models"
	"log"
	"net/http"
)

type EmailStruct struct {
}

type EmailBuilder interface{
	 CreateActionEmail([]byte)
	 CreateNotificationEmail([]byte)
	 CreateSubscriptionEmail([]byte)
}

func (c *EmailStruct) CreateActionEmail(body []byte) {

	p := models.IncomingActionDataStruct{}
	t.ToStruct(body, &p)

	template, subject := smtp.BaseTypeActionEmail(p.Template, p)

	go smtp.SendEmail([]string{p.Email}, subject, template)
	log.Printf("Email sent")

}

func (c *EmailStruct) CreateNotificationEmail(body []byte) {

	p := models.IncomingNotificationDataStruct{}
	t.ToStruct(body, &p)

	template, subject := smtp.BaseTypeNotificationEmail(p.Template, p)

	smtp.SendEmail([]string{p.Email}, subject, template)
	log.Printf("Email sent")

}

func (c *EmailStruct) CreateSubscriptionEmail(body []byte) {

	p := models.IncomingPaymentDataStruct{}
	t.ToStruct(body, &p)

	template, subject := smtp.BaseTypeSubscriptionEmail(p.Template, p)

	smtp.SendEmail([]string{p.Email}, subject, template)
	log.Printf("Email sent")

}

func CheckBodyStatus(w http.ResponseWriter, r *http.Request) {
	if r.ContentLength < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No body error"))
		return
	}
}

