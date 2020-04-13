package queue_types

import (
	t "github.com/ProjectReferral/Get-me-in/customer-service/models"
	"github.com/ProjectReferral/Get-me-in/customer-service/configs"
	"github.com/streadway/amqp"
)

func NotificationEmailQueue(s string, template string, queue string) {

	conn, err := amqp.Dial(configs.BrokerUrl)

	FailOnError(err, "Failed to connect to RabbitMQ")

	ReceiveAndProcess(s, conn, t.BASETYPE_NOTIFICATION, template, queue)
}