package event_driven

import (
	"github.com/ProjectReferral/Get-me-in/queueingt-api/configs"
	"github.com/ProjectReferral/Get-me-in/queueingt-api/internal/models"
	"github.com/streadway/amqp"
	"fmt"
	"log"
	"encoding/json"
	"strings"
	"net/http"
)

func TestQ(w http.ResponseWriter) bool{
	conn := createConnection(w)
	if conn == nil {
		log.Println("Error: Failed to connect to RabbitMQ")
		w.WriteHeader(http.StatusNotFound)
		return false
	}else{
		defer conn.Close()
	}
	return true
}

func RabbitCreateQueue(w http.ResponseWriter, queue models.QueueDeclare) {
	fmt.Printf("value is: %+v\n", queue)
	ch, conn := createChanel(w)
	if ch != nil {
		defer conn.Close()
		defer ch.Close()
		_, err := ch.QueueDeclare(
			queue.Name,
			queue.Durable,
			queue.DeleteWhenUnused,
			queue.Exclusive,
			queue.NoWait,
			queue.Arguments,
		)
		checkError(w,err)
	}
}

func RabbitCreateExchange(w http.ResponseWriter, exchange models.ExchangeDeclare) {
	fmt.Printf("value is: %+v\n", exchange)
	ch, conn := createChanel(w)
	if ch != nil {
		defer conn.Close()
		defer ch.Close()
		err := ch.ExchangeDeclare(
			exchange.Name,
			exchange.Type,
			exchange.Durable,
			exchange.DeleteWhenUnused,
			exchange.Exclusive,
			exchange.NoWait,
			exchange.Arguments,
		)
		checkError(w,err)
	}
}

func RabbitQueueBind(w http.ResponseWriter, bind models.QueueBind) {
	fmt.Printf("value is: %+v\n", bind)
	ch, conn := createChanel(w)
	if ch != nil {
		defer conn.Close()
		defer ch.Close()
		err := ch.QueueBind(
			bind.Name,
			bind.Key,
			bind.Exchange,
			bind.DeleteWhenUnused,
			bind.Arguments,
		)
		checkError(w,err)
	}
}

func RabbitPublish(w http.ResponseWriter, publish models.ExchangePublish) {
	fmt.Printf("value is: %+v\n", publish)
	ch, conn := createChanel(w)
	if ch != nil {
		defer conn.Close()
		defer ch.Close()
		err := ch.Publish(
			publish.Exchange,
			publish.Key,
			publish.Mandatory,
			publish.Immediate,
			publish.Publishing,
		)
		checkError(w,err)
	}
}

func RabbitConsume(w http.ResponseWriter, consume models.QueueConsume) {
	fmt.Printf("value is: %+v\n", consume)
	ch, conn := createChanel(w)
	if ch != nil {
		defer conn.Close()
		defer ch.Close()
		msgs, err := ch.Consume(
			consume.Name,
			consume.Consumer,
			consume.AutoAck,
			consume.Exclusive,
			consume.NoLocal,
			consume.NoWait,
			consume.Arguments,
        )
		if !checkError(w,err){
			var arr = []string{}
			var end bool
			for {
				fmt.Printf("messages len=%d cap=%d %v\n", len(arr), cap(arr), arr)
				select {
					case msg := <-msgs:
						body := string(msg.Body[:])
						log.Println("body is: ",body)
						arr = append(arr,body)
						end = false
					default:
						end = true
				}
				if(end){
					break
				}
			}
			json, readError := json.Marshal(arr)
			if !checkError(w,readError){
				log.Println("sending: ",string(json)," url: ",consume.URL)
				w.Write([]byte(string(json))) //TODO: implement subscription feature
			}else{
				log.Println("error converting to json\n")
			}
		}
	}
}

func createChanel(w http.ResponseWriter) (*amqp.Channel, *amqp.Connection) {
	conn := createConnection(w)
	if conn != nil {
		ch, err := conn.Channel()
		if !checkError(w,err) {
			return ch,conn
		}else{
			defer conn.Close()
		}
	}
	return nil,nil
}


func createConnection(w http.ResponseWriter) *amqp.Connection {
	conn, err := amqp.Dial(configs.BrokerUrl)
	if checkError(w,err) {
		return nil
	}
	return conn
}

func checkError(w http.ResponseWriter, err error) bool{
	if err != nil {
		log.Println(err.Error())
		if strings.Contains(err.Error(),"Exception (404)") {
			w.WriteHeader(404)
		}else{
			w.WriteHeader(400)
		}
		return true
	}
	return false
}