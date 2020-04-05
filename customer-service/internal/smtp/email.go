package smtp

import (
	"fmt"
	"log"
	"net/smtp"
	"html/template"
)

func SendEmail(to []string, subject string, body string) {

	auth := smtp.PlainAuth(
		"",
		"project181219@gmail.com",
		"",
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



	// func SendEmail(to []string, subject string, body string, templateFileName string) {

	// 	t, errr := template.ParseFiles(templateFileName)
	// 	if errr != nil {
	// 		log.Fatal(errr)
	// 	}

	// 	auth := smtp.PlainAuth(
	// 		"",
	// 		"project181219@gmail.com",
	// 		"",
	// 		"smtp.gmail.com",
	// 	)
	
	// 	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	
	// 	sub := fmt.Sprintf("Subject: %s\n", subject)
	// 	msg := []byte(sub + mime + fmt.Sprintf("<html><body><h1>%s</h1></body></html>", body))

	// 	err := smtp.SendMail(
	// 		"smtp.gmail.com:587",
	// 		auth,
	// 		"project181219@gmail.com",
	// 		to,
	// 		msg,
	// 	)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }


