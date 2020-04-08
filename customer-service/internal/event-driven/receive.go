package event_driven

import (
	"github.com/ProjectReferral/Get-me-in/customer-service/configs"
	s "github.com/ProjectReferral/Get-me-in/customer-service/models"
	"github.com/ProjectReferral/Get-me-in/notification-service/smtp"
	"github.com/streadway/amqp"
	"log"
)

func ReceiveFromAllQs() {
	conn, err := amqp.Dial(configs.BrokerUrl)

	failOnError(err, "Failed to connect to RabbitMQ")
	//failOnError(err, "Failed to declare a queue")

	ActionEmailType("new-user-verify-email", "Please verify your email", conn, s.ActionEmailStruct{
		Intro:       "Welcome to GMI! We're very excited to have you on board.",
		Instruct:    "To get started, please click here:",
		ButtonText:  "Confirm your account",
		ButtonColor: "#22BC66",
		Outro:       "Need help, or have questions? Just reply to this email, we'd love to help.",
	})

	// TODO -  WILL NEED TO FILL OUT WHEN RESET-USER-EMAIL Q IS MADE, THIS ALSO GOES FOR THE FUNCS BELOW!!!
	ActionEmailType("reset-user-email", "Reset your password", conn, s.ActionEmailStruct{
		Intro:       "------",
		Instruct:    "------",
		ButtonText:  "------",
		ButtonColor: "#fc2403",
		Outro:       "------",
	})

	NotificationEmailType("payment-confirmation-email", "Payment confirmation", conn, s.NotificationEmailStruct{
		Intro: "------",
		Outro: "------",
	})

	NotificationEmailType("cancel-subscription-email", "Canceled subscription", conn,s.NotificationEmailStruct{
		Intro: "------",
		Outro: "------",
	})

	//Debugging purposes
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
}

func ActionEmailType(queue string, subject string, conn *amqp.Connection, c s.ActionEmailStruct) {

	defer conn.Close()
	ch, err := conn.Channel()

	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	msgsCreateUser, err := ch.Consume(
		queue, // queue
		"",                    // consumer
		false,                  // auto-ack,
		false,                 // exclusive
		false,                 // no-local
		false,                 // no-wait
		nil,                     // args
	)

	failOnError(err, "Failed to register a consumer")

	forever := make(chan string)

	go func() {
		for d := range msgsCreateUser {

			smtp.SendEmail([]string{"sumite3117@hotmail.com", "sharjeel50@hotmail.co.uk", "hamza_razeq@hotmail.co.uk", "0101hamza@gmail.com"}, subject, smtp.ActionEmail(s.ActionEmailStruct{
				Name:        "Hamza",
				Intro:       c.Intro,
				Instruct:    c.Instruct,
				ButtonText:  c.ButtonText,
				ButtonColor: c.ButtonColor,
				ButtonLink:  "-",
				Outro:       c.Outro,
			}))

			log.Printf("Received a message: %s - %s", d.Body, d.CorrelationId)
			d.Ack(true)
		}
	}()

	<-forever
}

func NotificationEmailType(queue string, subject string, conn *amqp.Connection, c s.NotificationEmailStruct) {

	defer conn.Close()
	ch, err := conn.Channel()

	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	msgsCreateUser, err := ch.Consume(
		queue, // queue
		"",                    // consumer
		false,                  // auto-ack,
		false,                 // exclusive
		false,                 // no-local
		false,                 // no-wait
		nil,                     // args
	)

	failOnError(err, "Failed to register a consumer")

	forever := make(chan string)

	go func() {
		for d := range msgsCreateUser {

			// Will need to give this a body for more of dat blingbling
			smtp.SendEmail([]string{"sumite3117@hotmail.com"}, subject, smtp.NotificationEmail(s.NotificationEmailStruct{
				Name:  "a",
				Intro: c.Intro,
				Outro: c.Outro,
			}))

			log.Printf("Received a message: %s - %s", d.Body, d.CorrelationId)
			d.Ack(true)
		}
	}()

	<-forever
}



//func CreateChannel(conn *amqp.Connection, queue string) *amqp.Channel {
//	defer conn.Close()
//	ch, err := conn.Channel()
//
//	failOnError(err, "Failed to open a channel")
//	defer ch.Close()
//
//	msgsCreateUser, err := ch.Consume(
//		queue, // queue
//		"",                    // consumer
//		false,                  // auto-ack,
//		false,                 // exclusive
//		false,                 // no-local
//		false,                 // no-wait
//		nil,                     // args
//	)
//
//	failOnError(err, "Failed to register a consumer")
//}

