package smtp

import (
	"github.com/ProjectReferral/Get-me-in/customer-service/models"
	"github.com/matcornic/hermes"
	"log"
)


// Configure hermes by setting a theme and your product info
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

// This will be used for two types of emails currently, reset password and email confirmation.
func ActionEmail(params models.ActionEmailStruct) string {

	email := hermes.Email{
		Body: hermes.Body{
			Name: params.Name,
			Intros: []string {
				params.Intro,
			},
			Actions: []hermes.Action{
				{
					Instructions: params.Instruct,
					Button: hermes.Button{
						Color: params.ButtonColor,
						Text:  params.ButtonText,
						Link:  params.ButtonLink,
					},
				},
			},
			Outros: []string{
				params.Outro,
			},
		},
	}

	emailBody, err := h.GenerateHTML(email)
	if err != nil {
		failOnError(err, "Failed to generate HTML email")
	}
	return emailBody
}

// This email will be used to only notifiy a user without any actionss
func NotificationEmail(params models.NotificationEmailStruct) string {
	email := hermes.Email{
		Body: hermes.Body{
			Name: params.Name,
			Intros: []string{
				params.Intro,
			},
			Outros: []string{
				params.Outro,
			},
		},
	}

	emailBody, err := h.GenerateHTML(email)
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
