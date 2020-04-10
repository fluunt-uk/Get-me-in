package queue_types

import (
	"github.com/ProjectReferral/Get-me-in/customer-service/configs"
	"github.com/streadway/amqp"
)

func ActionEmailQueue(d string, s string, queue string) {

	conn, err := amqp.Dial(configs.BrokerUrl)

	FailOnError(err, "Failed to connect to RabbitMQ")

	ReceiveAndProcess(d, s, conn, "action", queue)
}

