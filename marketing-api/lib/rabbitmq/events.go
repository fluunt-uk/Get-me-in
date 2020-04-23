package rabbitmq

import (
	"github.com/ProjectReferral/Get-me-in/queueing-api/client"
	"github.com/ProjectReferral/Get-me-in/queueing-api/client/models"
	"github.com/streadway/amqp"
	"log"
	"net/http"
)

var Client client.QueueClient

func BroadcastNewAdvert(body []byte){

	client := &http.Client{}

	//not dependant on the response
	_, err := Client.Publish(client, models.ExchangePublish{
		Exchange:   "accounts.fanout",
		Key:        "",
		Mandatory:  false,
		Immediate:  false,
		Publishing: amqp.Publishing{
			ContentType:   "text/plain",
			Body:          body,
			CorrelationId: "correlationId",
		},
	})

	if err != nil {
		log.Printf("Http request to RabbitMQ API failed with :[%s]", err.Error())
	}
}