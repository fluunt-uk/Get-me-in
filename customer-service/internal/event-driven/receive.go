package event_driven

import (
	"encoding/json"
	"log"

	"github.com/ProjectReferral/Get-me-in/account-api/configs"
	email_configs "github.com/ProjectReferral/Get-me-in/customer-service/configs"
	"github.com/ProjectReferral/Get-me-in/customer-service/notification-service/smtp"
	"github.com/streadway/amqp"
)

func ReceiveFromAllQs() {
	conn, err := amqp.Dial(configs.BrokerUrl)

	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	failOnError(err, "Failed to declare a queue")

	msgsCreateUser, err := ch.Consume(
		"new-user-verify-email", // queue
		"",                      // consumer
		false,                   // auto-ack, TODO: manual ack
		false,                   // exclusive
		false,                   // no-local
		false,                   // no-wait
		nil,                     // args
	)

	failOnError(err, "Failed to register a consumer")

	forever := make(chan string)
	
	go func() {
		for d := range msgsCreateUser {

			smtp.SendEmail([]string{"sumite3117@hotmail.com"}, "Please verify your email", smtp.ActionEmail(email_configs.ActionEmail{Name: "",
				Intro: "Intro test",
				Instruct: "Please click on the button below to get started",
				ButtonText: "Verify Email",
				ButtonColor: "#22BC66",
				ButtonLink: "LIIIINKKKKK",
				Outro: "Thank you :)"}))

			log.Printf("Received a message: %s - %s", d.Body, d.CorrelationId)
			d.Ack(true)
		}
	}()

	<-forever

	//Debugging purposes
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
}
