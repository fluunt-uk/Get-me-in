package smtp

import (
	"fmt"
	"log"
	"net/smtp"
	"html/template"
	"github.com/matcornic/hermes"
)

func SendEmail(to []string, subject string, emailTemplate func) {

	t, errr := template.ParseFiles(emailTemplate)
	if errr != nil {
		log.Fatal(errr)
	}

	var body btyes.Buffer

	auth := smtp.PlainAuth(
				"",
				"project181219@gmail.com",
				"",
				"smtp.gmail.com",
			)

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	body.write([]byte(fmt.Sprintf("Subject: %s\n%s\n\n", subject, mime)))


	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"project181219@gmail.com",
		to,
		body.Bytes()
	)
	if err != nil {
		log.Fatal(err)
	}
}

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


