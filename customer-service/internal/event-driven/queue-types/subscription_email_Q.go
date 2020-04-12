package queue_types

import (
	"github.com/ProjectReferral/Get-me-in/customer-service/configs"
	"github.com/streadway/amqp"
)

func SubscriptionEmailQueue(s string, template string, queue string) {

	conn, err := amqp.Dial(configs.BrokerUrl)

	FailOnError(err, "Failed to connect to RabbitMQ")

	ReceiveAndProcess(s, conn, template, queue)
}


//import (
//	event_driven "github.com/ProjectReferral/Get-me-in/customer-service/internal/event-driven"
//	"github.com/ProjectReferral/Get-me-in/customer-service/internal/smtp"
//	"github.com/ProjectReferral/Get-me-in/customer-service/internal/templates"
//	s "github.com/ProjectReferral/Get-me-in/customer-service/models"
//	"github.com/streadway/amqp"
//	"log"
//)

//func PaymentEmailType(queue string, subject string, conn *amqp.Connection) {
//
//	defer conn.Close()
//	ch, err := conn.Channel()
//
//	FailOnError(err, "Failed to open a channel")
//	defer ch.Close()
//
//	msgsCreateUser, err := ch.Consume(
//		queue, // queue
//		"",
//		false,
//		false,
//		false,
//		false,
//		nil,
//	)
//
//	FailOnError(err, "Failed to register a consumer")
//
//	forever := make(chan string)
//
//	go func() {
//		for d := range msgsCreateUser {
//
//			o := event_driven.RetrieveData("payment", d)
//			k := o.(s.IncomingPaymentDataStruct)
//
//			// Will need to give this a body for more of dat blingbling
//			smtp.SendEmail([]string{k.Email}, subject, templates.PaymentEmail(s.PaymentEmailStruct{
//				Firstname:  k.Fullname(),
//				Premium: k.Premium,
//				Description: k.Description,
//				Price: k.Price,
//			}))
//
//			log.Printf("Received a message: %s - %s", d.Body, d.CorrelationId)
//			d.Ack(true)
//		}
//	}()
//
//	<-forever
//}
