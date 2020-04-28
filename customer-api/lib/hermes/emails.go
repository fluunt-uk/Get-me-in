package hermes

import (
	"fmt"
	"github.com/ProjectReferral/Get-me-in/customer-api/internal/smtp"
	t "github.com/ProjectReferral/Get-me-in/customer-api/lib/hermes/templates"
	"github.com/ProjectReferral/Get-me-in/customer-api/models"
	"io/ioutil"
	"log"
	"net/http"
)

type EmailStruct struct {
}

type EmailBuilder interface{
	 CreateActionEmail([]byte)
	 CreateNotificationEmail(http.ResponseWriter, *http.Request)
	 CreateSubscriptionEmail(http.ResponseWriter, *http.Request)
}

func (c *EmailStruct) CreateActionEmail(body []byte) {

	//CheckBodyStatus(w, r)
	//s, err := ioutil.ReadAll(r.Body)
	//
	//if(err != nil){
	//	fmt.Println(err)
	//}

	p := models.IncomingActionDataStruct{}
	//t.ToStruct(s, &p)

	template, subject := smtp.BaseTypeActionEmail(p.Template, p)

	go smtp.SendEmail([]string{p.Email}, subject, template)
	log.Printf("Email sent")

}

func (c *EmailStruct) CreateNotificationEmail(w http.ResponseWriter, r *http.Request) {

	CheckBodyStatus(w, r)

	s, err := ioutil.ReadAll(r.Body)

	if(err != nil){
		fmt.Println(err)
	}

	p := models.IncomingNotificationDataStruct{}
	t.ToStruct(s, &p)

	template, subject := smtp.BaseTypeNotificationEmail(p.Template, p)

	smtp.SendEmail([]string{p.Email}, subject, template)
	log.Printf("Email sent")

}

func (c *EmailStruct) CreateSubscriptionEmail(w http.ResponseWriter, r *http.Request) {

	CheckBodyStatus(w, r)

	s, err := ioutil.ReadAll(r.Body)

	if(err != nil){
		fmt.Println(err)
	}

	p := models.IncomingPaymentDataStruct{}
	t.ToStruct(s, &p)

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
