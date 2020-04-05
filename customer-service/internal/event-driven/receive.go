package event_driven

import (
	"log"

	"github.com/ProjectReferral/Get-me-in/account-api/configs"
	"github.com/ProjectReferral/Get-me-in/customer-service/internal/smtp"
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
	//"hamza_razeq@hotmail.co.uk", "0101hamza@gmail.com",
	go func() {
		// Will ne
		for d := range msgsCreateUser {

			smtp.SendEmail(
				[]string{"sharjeel50@hotmail.co.uk", "sumite3117@hotmail.com", "hamza_razeq@hotmail.co.uk", "0101hamza@gmail.com"},
				"Verify your email - TESTING",
				"Please verify your email - TESTING")

			log.Printf("Received a message: %s - %s", d.Body, d.CorrelationId)

			d.Ack(true)
		}
	}()

	<-forever

	//Debugging purposes
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
}
