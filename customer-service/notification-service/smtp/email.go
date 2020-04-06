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
				"Newproject2019",
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


