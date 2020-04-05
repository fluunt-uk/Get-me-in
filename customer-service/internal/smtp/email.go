package smtp

import (
	"fmt"
	"log"
	"net/smtp"
)

func SendEmail(to []string, subject string, body string) {

	auth := smtp.PlainAuth(
		"",
		"project181219@gmail.com",
		"Newproject2019",
		"smtp.gmail.com",
	)

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	sub := fmt.Sprintf("Subject: %s\n", subject)
	msg := []byte(sub + mime + fmt.Sprintf("<html><body><h1>%s</h1></body></html>", body))

	// msg := []byte(fmt.Sprintf("Subject: %s\r\n"+
	// 	"\r\n"+
	// 	"%s\r\n", subject, body))

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"project181219@gmail.com",
		to,
		msg,
	)
	if err != nil {
		log.Fatal(err)
	}
}
