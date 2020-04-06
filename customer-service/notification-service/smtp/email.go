package smtp

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"html/template"
	"github.com/matcornic/hermes"
)

func SendEmail(to []string, subject string, html_template string) {

	t, errr := template.ParseFiles(html_template)
	if errr != nil {
		log.Fatal(errr)
	}

	var body bytes.Buffer

	auth := smtp.PlainAuth(
				"",
				"project181219@gmail.com",
				"Newproject2019",
				"smtp.gmail.com",
			)

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	body.Write([]byte(fmt.Sprintf("Subject: %s\n%s\n\n", subject, mime)))


	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"project181219@gmail.com",
		to,
		body.Bytes(),
	)
	if err != nil {
		log.Fatal(err)
	}
}


