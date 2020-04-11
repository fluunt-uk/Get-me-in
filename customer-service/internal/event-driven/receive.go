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

	queue_types.ActionEmailQueue("Please verify your email", "new-user","new-user-verify-email")
	queue_types.ActionEmailQueue("Reset your password", "reset-password","reset-password-email")
	queue_types.NotificationEmailQueue("Canceled subscription", "cancel-subscription", "cancel-subscription-email")

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