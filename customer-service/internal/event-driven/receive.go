package event_driven

import (
	"encoding/json"
	"fmt"
	"github.com/ProjectReferral/Get-me-in/customer-service/configs"
	s "github.com/ProjectReferral/Get-me-in/customer-service/models"
	"github.com/ProjectReferral/Get-me-in/customer-service/notification-service/smtp"
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

	//ActionEmailType("reset-user-email", "Reset your password", conn, s.ActionEmailStruct{
	//	Intro:       "------",
	//	Instruct:    "------",
	//	ButtonText:  "------",
	//	ButtonColor: "#fc2403",
	//	Outro:       "------",
	//})
	//
	//NotificationEmailType("payment-confirmation-email", "Payment confirmation", conn, s.NotificationEmailStruct{
	//	Intro: "------",
	//	Outro: "------",
	//})
	//
	//NotificationEmailType("cancel-subscription-email", "Canceled subscription", conn, s.NotificationEmailStruct{
	//	Intro: "------",
	//	Outro: "------",
	//})
	//
	//NotificationEmailType("create-subscription-email", "Subscription successfully created", conn, s.NotificationEmailStruct{
	//	Intro: "------",
	//	Outro: "------",
	//})
	//
	//NotificationEmailType("reminder-email", "Reminder", conn, s.NotificationEmailStruct{
	//	Intro: "------",
	//	Outro: "------",
	//})
	//
	//PaymentEmailType("payment-notification-email", "Subscription confirmation", conn)


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

			//var p s.IncomingDataStruct
			//err := json.Unmarshal(d.Body, &p)
			//if err != nil {
			//	fmt.Println(err)
			//}
			//fmt.Println(p)

			i := RetrieveData("action", d)

			// , "hamza_razeq@hotmail.co.uk", "0101hamza@gmail.com", "sumite3117@hotmail.com"
			smtp.SendEmail([]string{i.Email}, subject, smtp.ActionEmail(s.ActionEmailStruct{
				Name:        i.Firstname + i.Surname,
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
		"",
		false,
		false,
		false,
		false,
		nil,
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

func PaymentEmailType(queue string, subject string, conn *amqp.Connection) {

	defer conn.Close()
	ch, err := conn.Channel()

	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	msgsCreateUser, err := ch.Consume(
		queue, // queue
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	failOnError(err, "Failed to register a consumer")

	forever := make(chan string)

	go func() {
		for d := range msgsCreateUser {

			// Will need to give this a body for more of dat blingbling
			smtp.SendEmail([]string{"sumite3117@hotmail.com"}, subject, smtp.PaymentEmail(s.PaymentEmailStruct{
				Firstname:  "--",
				Premium: "--",
				Description: "--",
				Price: "--",
			}))

			log.Printf("Received a message: %s - %s", d.Body, d.CorrelationId)
			d.Ack(true)
		}
	}()

	<-forever
}

func RetrieveData(typeof string, d amqp.Delivery) interface{} {

	var p interface{}

	switch typeof {
	case "notification":
		p = s.IncomingNotificationDataStruct{}
	case "payment":
		p = s.IncomingPaymentDataStruct{}
	case "action":
		p = s.IncomingActionDataStruct{}
	}

	err := json.Unmarshal(d.Body, &p)
	if err != nil {
		fmt.Println(err)
	}

	return p
}

//func CreateChannel(conn *amqp.Connection, queue string) *amqp.Channel {
//
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

