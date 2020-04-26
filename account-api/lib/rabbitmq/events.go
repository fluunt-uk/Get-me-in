package rabbitmq

import (
	"crypto/rand"
	"fmt"
	"github.com/ProjectReferral/Get-me-in/account-api/configs"
	"github.com/ProjectReferral/Get-me-in/queueing-api/client"
	"github.com/ProjectReferral/Get-me-in/queueing-api/client/models"
	"github.com/streadway/amqp"
	"log"
	"net/http"
)

var Client client.QueueClient

func BroadcastUserCreatedEvent(body []byte) {

	client := &http.Client{}

	//not dependant on the response
	_, err := Client.Publish(client, models.ExchangePublish{
		Exchange:  configs.FANOUT_EXCHANGE,
		Key:       "",
		Mandatory: false,
		Immediate: false,
		Publishing: amqp.Publishing{
			ContentType:   "text/plain",
			Body:          body,
			CorrelationId: NewUUID(),
		},
	})

	if err != nil {
		log.Printf("Http request to RabbitMQ API failed with :[%s]", err.Error())
	}
}

func NewUUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	return uuid
}
