package smtp

import 	"github.com/matcornic/hermes"

// This will be used for two types of emails currently, reset password and email confirmation.
func ActionEmail(name string, intro string, instruc string, buttonText string, buttonColor string, buttonLink string, outro string) hermes.Email {

	email := hermes.Email{
		Body: hermes.Body{
			Name: name,
			Intros: []string{
				intro,
			},
			Actions: []hermes.Action{
				{
					Instructions: instruc,
					Button: hermes.Button{
						Color: buttonColor,
						Text:  buttonText,
						Link:  buttonLink,
					},
				},
			},
			Outros: []string{
				outro,
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
func NotificationEmail(name string, intro string, outro string) hermes.Email {
	email := hermes.Email{
		Body: hermes.Body{
			Name: name,
			Intros: []string{
				intro,
			},
			Outros: []string{
				outro,
			},
		},
	}

	emailBody, err := h.GenerateHTML(email)
	if err != nil {
		failOnError(err, "Failed to generate HTML email")
	}
	
	return emailBody
}

// TODO
func PaymentsEmail() hermes.Email {

	emailBody, err := h.GenerateHTML(email)
	if err != nil {
		failOnError(err, "Failed to generate HTML email")
	}
	
	return emailBody
}