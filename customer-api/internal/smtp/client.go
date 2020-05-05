package smtp

import (
	"github.com/ProjectReferral/Get-me-in/customer-api/configs"
	"log"
	"net/smtp"
)

func SendEmail(to []string, subject string, html_template string) {

	auth := smtp.PlainAuth(
				"",
				configs.DevEmail,
				configs.DevEmailPw,
				"smtp.gmail.com",
			)

	msg := []byte("Subject:" + subject + "\n" +"MIME-version: 1.0;\nContent-Type: text/html; " +
		"charset=\"UTF-8\";\n\n" + html_template)

	err := smtp.SendMail("smtp.gmail.com:587", auth,
		configs.DevEmail, to, msg,)

	if err != nil {
		log.Println(err)
	}

	log.Printf("Email sent")
}


