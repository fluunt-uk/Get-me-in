package queue_types

import (
	"github.com/ProjectReferral/Get-me-in/customer-service/configs"
	"github.com/streadway/amqp"
)

func NotificationEmailQueue(s string, template string, queue string) {

	conn, err := amqp.Dial(configs.BrokerUrl)

	FailOnError(err, "Failed to connect to RabbitMQ")

	ReceiveAndProcess(s, conn, template, queue)
}