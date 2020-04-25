package queue_types

import (
	"github.com/ProjectReferral/Get-me-in/customer-service/internal/smtp"
	"github.com/ProjectReferral/Get-me-in/customer-service/lib/hermes/templates"
	"github.com/streadway/amqp"
	"log"
)

func ReceiveAndProcess(subject string, basetype string, template string, queue string){


			template, email := templates.GenerateHTMLTemplate(basetype, template, d.Body)

			smtp.SendEmail([]string{email}, subject, template)
			log.Printf("Email sent")
}