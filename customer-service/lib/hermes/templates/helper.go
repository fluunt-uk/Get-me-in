package templates

import (
	"encoding/json"
	"fmt"
	s "github.com/ProjectReferral/Get-me-in/customer-service/models"
	"github.com/matcornic/hermes"
	"log"
)


// Configure templates by setting a theme and your product info
var h = hermes.Hermes{
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
}

func GenerateHTMLTemplate(typeof string, d []byte) (string, string) {

	switch typeof {
	case "cancel-subscription":
		p := s.IncomingNotificationDataStruct{}

		toStruct(d, &p)

		return GenerateNotificationHTMLTemplate(p, s.NotificationEmailStruct{
			Intro: "",
			Outro: "",
		})

	case "":
		p := s.IncomingNotificationDataStruct{}

		toStruct(d, &p)

		return GenerateNotificationHTMLTemplate(p, s.NotificationEmailStruct{
			Intro: "",
			Outro: "",
		})

	case "payment":
		p := s.IncomingPaymentDataStruct{}

		toStruct(d, &p)
		return GenerateSubscriptionHTMLTemplate(p, s.PaymentEmailStruct{})

	case "new-user":
		p := s.IncomingActionDataStruct{}

		toStruct(d, &p)

		return GenerateActionHTMLTemplate(p, s.ActionEmailStruct{
			Intro:       "Welcome to GMI! We're very excited to have you on board.",
			Instruct:    "To get started, please click here:",
			ButtonText:  "Confirm your account",
			ButtonColor: "#22BC66",
			Outro:       "Need help, or have questions? Just reply to this email, we'd love to help.",
		})

	case "reset-password":
		p := s.IncomingActionDataStruct{}

		toStruct(d, &p)

		return GenerateActionHTMLTemplate(p, s.ActionEmailStruct{
			Intro:       "You recently made a request to reset your password.",
			Instruct:    "Please click the link below to continue.",
			ButtonText:  "Reset Password",
			ButtonColor: "#fc2403",
			Outro:       "If you did not make this change or you believe an unauthorised person has accessed your account, go to {reset-password endpoint} to reset your password without delay.",
		})
	}

	return "", ""
}

func toStruct(d []byte, p interface{}){
	err := json.Unmarshal(d, &p)
	if err != nil {
		fmt.Println(err)
	}
}

func StringParsedHTML(e hermes.Email) string {
	emailBody, err := h.GenerateHTML(e)
	if err != nil {
		failOnError(err, "Failed to generate HTML email")
	}
	return emailBody
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}