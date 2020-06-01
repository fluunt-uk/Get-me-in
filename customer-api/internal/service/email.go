package service

import (
	"github.com/ProjectReferral/Get-me-in/customer-api/configs"
	"github.com/ProjectReferral/Get-me-in/customer-api/internal/smtp"
	t "github.com/ProjectReferral/Get-me-in/customer-api/lib/hermes/templates"
	"github.com/ProjectReferral/Get-me-in/customer-api/models"
	"github.com/matcornic/hermes"
)

type EmailService struct {
	EB *t.EmailBuilder
}

func (c *EmailService) SendEmail(body []byte) {

	p := models.IncomingData{}
	t.ToStruct(body, &p)

	template, subj := c.EB.GenerateHTMLTemplate(p)

	go smtp.SendEmail([]string{p.Email}, subj, template)
}

//all templates added to our map
func (c *EmailService) SetupTemplates(){
	c.EB.Init()

	c.EB.SetTheme(	&hermes.Hermes{
		// Optional Theme
		// Theme: new(Default)
		Product: hermes.Product{
			// Appears in header & footer of e-mails
			Name: "GMI Team",
			// Optional product logo
			Logo: "https://www.clipartmax.com/png/middle/425-4252869_blank-raffle-tickets-template-free-ticket-booking-icon-png.png",
			Copyright: "Copyright © 2020 GMI. All rights reserved.",
			TroubleText: "",
		},
	})

	c.EB.AddStaticTemplate(configs.NEW_USER_VERIFY,
		&models.BaseEmail{
			Subject: "----------",
			Intro: "Welcome to GMI! We're very excited to have you on board.",
			Outro: "Need help, or have questions? Just reply to this email, we'd love to help.",
			Action: models.ActionEmail{
				Instruct:    "To get started, please click here:",
				ButtonText:  "Confirm your account",
				ButtonColor: "#22BC66",
			},
		},
	)

	c.EB.AddStaticTemplate(configs.RESET_PASSWORD,
		&models.BaseEmail{
			Subject: "----------",
			Intro:       "You recently made a request to reset your password.",
			Outro:       "If you did not make this change or you believe an unauthorised person has accessed your account, go to {reset-password endpoint} to reset your password without delay.",
			Action: models.ActionEmail{
				Instruct:    "Please click the link below to continue.",
				ButtonText:  "Reset Password",
				ButtonColor: "#fc2403",
			},
		},
	)

	c.EB.AddStaticTemplate(configs.CREATE_SUBSCRIPTION,
		&models.BaseEmail{
			Subject: "----------",
			Intro: "Welcome! Your GMI experience just got premium.",
			Outro: "",
		},
	)

	c.EB.AddStaticTemplate(configs.CANCEL_SUBSCRIPTION,
		&models.BaseEmail{
			Subject: "----------",
			Intro: "This is a confirmation that your GMI account has been canceled at your request.",
			Outro: "To start applying again, you can reactivate your account at any time. We hope you decide to come back soon.",
		},
	)

	c.EB.AddStaticTemplate(configs.CANCEL_SUBSCRIPTION,
		&models.BaseEmail{
			Subject: "----------",
			Intro: "This is a confirmation that your GMI account has been canceled at your request.",
			Outro: "To start applying again, you can reactivate your account at any time. We hope you decide to come back soon.",
		},
	)

	c.EB.AddStaticTemplate(configs.REFEREE_APPLICATION,
		&models.BaseEmail{
		Subject: "----------",
		Intro: "",
			Outro: "",
		},
	)

	c.EB.AddStaticTemplate(configs.REMINDER,
		&models.BaseEmail{
			Subject: "----------",
			Intro: "",
			Outro: "",
		},
	)

	c.EB.AddStaticTemplate(configs.PAYMENT_CONFIRMATION,
		&models.BaseEmail{
			Subject: "----------",
			Intro: "Your order has been processed successfully.",
			Outro: "Thank you, enjoy your experience.",
		},
	)
	c.EB.AddStaticTemplate(configs.PAYMENT_INVOICE,
		&models.BaseEmail{
			Subject: "----------",
			Intro: "Your order has been processed successfully.",
			Outro: "Thank you, enjoy your experience.",
		},
	)
}

//func checkBodyStatus(w http.ResponseWriter, r *http.Request) {
//	if r.ContentLength < 1 {
//		w.WriteHeader(http.StatusBadRequest)
//		w.Write([]byte("No body error"))
//		return
//	}
//}