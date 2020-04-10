package event_driven

import (
	"github.com/ProjectReferral/Get-me-in/customer-service/internal/event-driven/queue-types"
	"log"
)

func ReceiveFromAllQs() {

	//failOnError(err, "Failed to declare a queue")

/*	queue_types.ActionEmailType("new-user-verify-email", "Please verify your email", conn, s.ActionEmailStruct{
		Intro:       "Welcome to GMI! We're very excited to have you on board.",
		Instruct:    "To get started, please click here:",
		ButtonText:  "Confirm your account",
		ButtonColor: "#22BC66",
		Outro:       "Need help, or have questions? Just reply to this email, we'd love to help.",
	})*/

	queue_types.ActionEmailQueue("jipesh14@gmail.com", "Please verify your email", "new-user-verify-email")

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


//life
//alien covenant
