package event_driven

import (
	"log"
	"encoding/json"

	"github.com/ProjectReferral/Get-me-in/account-api/configs"
	"github.com/ProjectReferral/Get-me-in/customer-service/notification-service/smtp"
	"github.com/ProjectReferral/Get-me-in/customer-service/configs"
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

	type User struct {
        Name  string
        Order string
    }

	go func() {
		// Will ne
		for d := range msgsCreateUser {

			err := json.Unmarshal(User, &animals)
    		if err != nil {
				fmt.Println("error:", err)
			}
			
			// name string, intro string, instruc string, buttonText string, buttonColor string, buttonLink string, outro string
			smtp.SendEmail("to", "subject", smtp.ActionEmail("name", 
															"Welcome to Get-me-in",
															"To get started applying, please click here:",
															"Verify Email",
															"#22BC66",
															"LinkTing",
															"Safesafesafe"))

			log.Printf("Received a message: %s - %s", d.Body, d.CorrelationId)
			d.Ack(true)
		}
	}()

	<-forever

	//Debugging purposes
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
}
