package queue_types

import (
	"github.com/ProjectReferral/Get-me-in/customer-service/configs"
	t "github.com/ProjectReferral/Get-me-in/customer-service/models"
	"github.com/streadway/amqp"
)

func ActionEmailQueue(s string, template string, queue string) {

	conn, err := amqp.Dial(configs.BrokerUrl)

	FailOnError(err, "Failed to connect to RabbitMQ")

	ReceiveAndProcess(s, conn, t.BASETYPE_ACTION, template, queue)
}

