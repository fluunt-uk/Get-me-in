package smtp

import (
	"log"
	"net/smtp"
)

//t, errr := template.ParseFiles(html_template)
//if errr != nil {
//	log.Fatal(errr)
//}
//body.Write([]byte(fmt.Sprintf("Subject: %s\n%s\n\n", subject, mime)))

func SendEmail(to []string, subject string, html_template string) {

	auth := smtp.PlainAuth(
				"",
				"project181219@gmail.com",
				"",
				"smtp.gmail.com",
			)

	msg := []byte("Subject:" + subject + "\n" +"MIME-version: 1.0;\nContent-Type: text/html; " +
		"charset=\"UTF-8\";\n\n" + html_template)

	err := smtp.SendMail("smtp.gmail.com:587", auth,
		"project181219@gmail.com", to, msg,)

	if err != nil {
		log.Fatal(err)
	}
}


