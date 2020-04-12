package queue_types

import (
	"github.com/ProjectReferral/Get-me-in/customer-service/internal/smtp"
	"github.com/ProjectReferral/Get-me-in/customer-service/lib/hermes/templates"
	"github.com/streadway/amqp"
	"log"
)

func ReceiveAndProcess(subject string, conn *amqp.Connection, template string, queue string){

	defer conn.Close()
	ch, err := conn.Channel()

	FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	msgsCreateUser, err := ch.Consume(
		queue,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	FailOnError(err, "Failed to register a consumer")

	forever := make(chan string)

	go func() {
		for d := range msgsCreateUser {
			log.Printf("Received a message: %s - %s", d.Body, d.CorrelationId)

			template, email := templates.GenerateHTMLTemplate(template, d.Body)

			smtp.SendEmail([]string{email}, subject, template)
			log.Printf("Email sent")
			d.Ack(true)
		}
	}()

	<-forever
}