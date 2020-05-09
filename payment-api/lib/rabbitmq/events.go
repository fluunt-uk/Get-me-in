package rabbitmq

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/ProjectReferral/Get-me-in/payment-api/configs"
	resource_model "github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api/resources/models"
	"github.com/ProjectReferral/Get-me-in/queueing-api/client"
	qm "github.com/ProjectReferral/Get-me-in/queueing-api/client/models"
	"github.com/streadway/amqp"
	"log"
	"net/http"
)

var Client client.QueueClient

func BroadcastNewSubEvent(s resource_model.Subscription) {

	client := &http.Client{}

	s.SetTemplate(configs.CREATE_SUBSCRIPTION)
	b, _ := json.Marshal(s)

	//not dependant on the response
	_, err := Client.Publish(client, qm.ExchangePublish{
		Exchange:  configs.FANOUT_EXCHANGE,
		Key:       "",
		Mandatory: false,
		Immediate: false,
		Publishing: amqp.Publishing{
			ContentType:   "text/plain",
			Body:          b,
			CorrelationId: NewUUID(),
		},
	})

	if err != nil {
		log.Printf("Http request to RabbitMQ API failed with :[%s]", err.Error())
	}

	log.Println("Message sent")
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
