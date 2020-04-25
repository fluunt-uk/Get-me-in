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

// Will return template, email, subject
func GenerateHTMLTemplate(basetype string, typeof string, d []byte) (string, string, string) {

	switch basetype {

	case s.BASETYPE_ACTION:
		return BaseTypeActionEmail(typeof, d)

	case s.BASETYPE_NOTIFICATION:
		return BaseTypeNotificationEmail(typeof, d)

	case s.BASETYPE_SUBSCRIPTION:
		return BaseTypeSubscriptionEmail(typeof, d)
	}

	return "","",""
}

func BaseTypeActionEmail(typeof string, d []byte) (string, string, string) {

	switch typeof {

	case s.NEW_USER_VERIFY:

		p := s.IncomingActionDataStruct{}
		toStruct(d, &p)

		t, e := GenerateActionHTMLTemplate(p, s.ActionEmailStruct{
			Intro:       "Welcome to GMI! We're very excited to have you on board.",
			Instruct:    "To get started, please click here:",
			ButtonText:  "Confirm your account",
			ButtonColor: "#22BC66",
			Outro:       "Need help, or have questions? Just reply to this email, we'd love to help.",
		})

		return t, e, ""

	case s.RESET_PASSWORD:

		p := s.IncomingActionDataStruct{}
		toStruct(d, &p)

		t, e := GenerateActionHTMLTemplate(p, s.ActionEmailStruct{
			Intro:       "You recently made a request to reset your password.",
			Instruct:    "Please click the link below to continue.",
			ButtonText:  "Reset Password",
			ButtonColor: "#fc2403",
			Outro:       "If you did not make this change or you believe an unauthorised person has accessed your account, go to {reset-password endpoint} to reset your password without delay.",
		})

		return t, e, ""
	}

	return "","",""
}

func BaseTypeNotificationEmail(typeof string, d []byte) (string, string, string) {

	switch typeof {

	case s.CREATE_SUBSCRIPTION:

		p := s.IncomingNotificationDataStruct{}
		toStruct(d, &p)

		t, e := GenerateNotificationHTMLTemplate(p, s.NotificationEmailStruct{
			Intro: "Welcome! Your GMI experience just got premium.",
			Outro: "",
		})

		return t, e, ""

	case s.CANCEL_SUBSCRIPTION:

		p := s.IncomingNotificationDataStruct{}
		toStruct(d, &p)

		// Will need to pass through some button to link to reactivate account (if possible) or pass in a button for sign up
		// Will also need to pass through when the service ends for the customer
		t, e := GenerateNotificationHTMLTemplate(p, s.NotificationEmailStruct{
			Intro: "This is a confirmation that your GMI account has been canceled at your request.",
			Outro: "To start applying again, you can reactivate your account at any time. We hope you decide to come back soon.",
		})

		return t, e, ""

	case s.REFEREE_APPLICATION:

		p := s.IncomingNotificationDataStruct{}
		toStruct(d, &p)

		// Will need to pass through some button to link to reactivate account (if possible) or pass in a button for sign up
		// Will also need to pass through when the service ends for the customer
		t, e := GenerateNotificationHTMLTemplate(p, s.NotificationEmailStruct{
			Intro: "",
			Outro: "",
		})

		return t, e, ""

	case s.REMINDER:

		p := s.IncomingNotificationDataStruct{}
		toStruct(d, &p)

		t, e := GenerateNotificationHTMLTemplate(p, s.NotificationEmailStruct{
			Intro: "",
			Outro: "",
		})

		return t, e, ""
	}

	return "","",""
}

func BaseTypeSubscriptionEmail(typeof string, d []byte) (string, string, string) {

	switch typeof {

	case s.PAYMENT_CONFIRMATION:

		p := s.IncomingPaymentDataStruct{}
		toStruct(d, &p)

		template, email := GenerateSubscriptionHTMLTemplate(p, s.PaymentEmailStruct{
			Intro: "Your order has been processed successfully.",
			Outro: "Thank you, enjoy your experience.",
		})

		return template, email, ""


	case s.PAYMENT_INVOICE:

		p := s.IncomingPaymentDataStruct{}
		toStruct(d, &p)

		t, e := GenerateSubscriptionHTMLTemplate(p, s.PaymentEmailStruct{})

		return t, e, ""
	}

	return "","",""
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