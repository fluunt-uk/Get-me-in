package event_driven

import (
	queue_types "github.com/ProjectReferral/Get-me-in/customer-service/internal/event-driven/queue-types"
	s "github.com/ProjectReferral/Get-me-in/customer-service/models"
	"log"
)

func ReceiveFromAllQs() {

	//queue_types.ActionEmailQueue("Please verify your email", s.NEW_USER_VERIFY,"new-user-verify-email")
	//queue_types.ActionEmailQueue("Reset your password", s.RESET_PASSWORD,"reset-password-email")
	//queue_types.SubscriptionEmailQueue("payment ting", s.PAYMENT_CONFIRMATION, "create-subscription-email")
	queue_types.NotificationEmailQueue("Canceled subscription", s.CANCEL_SUBSCRIPTION, "cancel-subscription-email")
	//queue_types.NotificationEmailQueue("Application", s.REFEREE_APPLICATION, "referee-application-email")

	//Debugging purposes
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
}