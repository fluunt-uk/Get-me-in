package smtp

import (
	"github.com/ProjectReferral/Get-me-in/customer-service/configs"
	"github.com/matcornic/hermes"
	event_driven "github.com/ProjectReferral/Get-me-in/customer-service/internal/event-driven"
)


// Configure hermes by setting a theme and your product info
h := hermes.Hermes{
    // Optional Theme
    // Theme: new(Default) 
    Product: hermes.Product{
        // Appears in header & footer of e-mails
        Name: "Hermes",
        Link: "https://example-hermes.com/",
        // Optional product logo
        Logo: "http://www.duchess-france.org/wp-content/uploads/2016/01/gopher.png",
    },
}

// This will be used for two types of emails currently, reset password and email confirmation.
func ActionEmail(params configs.ActionEmail) string {

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
		event_driven.failOnError(err, "Failed to generate HTML email")
	}
	return emailBody
}

// This email will be used to only notifiy a user without any actionss
func NotificationEmail(params configs.NotificationEmail) string {
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
		event_driven.failOnError(err, "Failed to generate HTML email")
	}
	
	return emailBody
}

// TODO
//func PaymentsEmail() hermes.Email {
//
//	emailBody, err := h.GenerateHTML(email)
//	if err != nil {
//		failOnError(err, "Failed to generate HTML email")
//	}
//
//	return emailBody
//}