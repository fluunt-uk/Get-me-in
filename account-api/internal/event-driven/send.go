package event_driven

import (
	"crypto/rand"
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type RabbitClient struct {
	URL string
	C *amqp.Connection
}

func (rc *RabbitClient) Connect (){
	conn, err := amqp.Dial(rc.URL)
	failOnError(err, "Failed to connect to RabbitMQ")
	rc.C = conn
}


func (rc *RabbitClient) SendToQ(body string, exchange string, correlationId string) {

	defer rc.C.Close()

	ch, err := rc.C.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	failOnError(err, "Failed to declare a queue")

	err = ch.Publish(
		exchange, // exchange
		"",       // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			Body:          []byte(body),
			CorrelationId: correlationId,
		})

	log.Println("Message sent:" + body)
	failOnError(err, "Failed to publish a message")
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

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
