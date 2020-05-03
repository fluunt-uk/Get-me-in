package service

import (
	"github.com/ProjectReferral/Get-me-in/customer-api/internal/smtp"
	t "github.com/ProjectReferral/Get-me-in/customer-api/lib/hermes/templates"
	"github.com/ProjectReferral/Get-me-in/customer-api/models"
)

type EmailService struct {
	AEB *t.EmailBuilder
}

func (c *EmailService) SendEmail(body []byte) {

	c.setupTemplates()
	p := models.IncomingData{}
	t.ToStruct(body, &p)

	template := c.AEB.GenerateHTMLTemplate(p)

	go smtp.SendEmail([]string{p.Email}, "Subject goes here", template)
}

//
//func (c *EmailService) CreateNotificationEmail(body []byte) {
//
//	p := models.IncomingData{}
//	t.ToStruct(body, &p)
//
//	template, subject := t.BaseTypeNotificationEmail(p.Template, p)
//
//	smtp.SendEmail([]string{p.Email}, subject, template)
//	log.Printf("Email sent")
//
//}
//
//func (c *EmailService) CreateSubscriptionEmail(body []byte) {
//
//	p := models.IncomingData{}
//	t.ToStruct(body, &p)
//
//	template, subject := t.BaseTypeSubscriptionEmail(p.Template, p)
//
//	smtp.SendEmail([]string{p.Email}, subject, template)
//	log.Printf("Email sent")
//
//}
//
//func checkBodyStatus(w http.ResponseWriter, r *http.Request) {
//	if r.ContentLength < 1 {
//		w.WriteHeader(http.StatusBadRequest)
//		w.Write([]byte("No body error"))
//		return
//	}
//}

func (c *EmailService) setupTemplates(){
	c.AEB.Innit()

	c.AEB.AddStaticTemplate(models.NEW_USER_VERIFY,
		&models.BaseEmail{
			Intro: "Welcome to GMI! We're very excited to have you on board.",
			Outro: "Need help, or have questions? Just reply to this email, we'd love to help.",
			Action: models.ActionEmail{
				Instruct:    "To get started, please click here:",
				ButtonText:  "Confirm your account",
				ButtonColor: "#22BC66",
			},
		},
	)

	c.AEB.AddStaticTemplate(models.RESET_PASSWORD,
		&models.BaseEmail{
			Intro:       "You recently made a request to reset your password.",
			Outro:       "If you did not make this change or you believe an unauthorised person has accessed your account, go to {reset-password endpoint} to reset your password without delay.",
			Action: models.ActionEmail{
				Instruct:    "Please click the link below to continue.",
				ButtonText:  "Reset Password",
				ButtonColor: "#fc2403",
			},
		},
	)
}