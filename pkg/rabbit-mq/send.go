package rabbit_mq

import (
	"crypto/rand"
	"fmt"
	"github.com/ProjectReferral/Get-me-in/pkg/rabbit-mq/configs"
	"github.com/streadway/amqp"
	"log"
)

func SendTest(routingKey string, body string, qName string, exchange string) {
	conn, err := amqp.Dial(configs.BrokerUrl)

	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	failOnError(err, "Failed to declare a queue")

	for i := 0; i < 5; i++ {
		err = ch.Publish(
			exchange,   // exchange
			routingKey, // routing key
			false,      // mandatory
			false,      // immediate
			amqp.Publishing{
				ContentType:   "text/plain",
				Body:          []byte(body),
				CorrelationId: newUUID(),
			})
	}

	for i := 0; i < 5; i++ {
		err = ch.Publish(
			exchange, // exchange
			"test1",  // routing key
			false,    // mandatory
			false,    // immediate
			amqp.Publishing{
				ContentType:   "text/plain",
				Body:          []byte("test from Queue1"),
				CorrelationId: newUUID(),
			})
	}

	failOnError(err, "Failed to publish a message")
}

func SendToQ(routingKey string, body string, qName string, exchange string) {
	conn, err := amqp.Dial(configs.BrokerUrl)

	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	failOnError(err, "Failed to declare a queue")

	err = ch.Publish(
		exchange,   // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	failOnError(err, "Failed to publish a message")
}

func newUUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	return uuid
}
