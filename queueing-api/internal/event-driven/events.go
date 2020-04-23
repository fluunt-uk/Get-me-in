package event_driven

import (
	"github.com/ProjectReferral/Get-me-in/queueing-api/configs"
	"github.com/ProjectReferral/Get-me-in/queueing-api/client/models"
	"github.com/ProjectReferral/Get-me-in/util"
	"github.com/streadway/amqp"
	"log"
	"encoding/json"
	"strings"
	"net/http"
	"time"
	"bytes"
	"os"
)

var consumers = map[string](chan bool){}
var messages  = map[uint64](uint64){}

type QueueMessage struct {
	ID       string  `json:"id"`
	Message  []byte  `json:"body"`
}

func TestQ(w http.ResponseWriter) bool{
	conn, err := amqp.Dial(configs.BrokerUrl)
	if conn == nil {
		log.Println(err.Error())
		log.Println("Error: Failed to connect to RabbitMQ")
		w.WriteHeader(http.StatusNotFound)
		return false
	}else{
		defer conn.Close()
	}
	return true
}

func RabbitCreateQueue(w http.ResponseWriter, queue models.QueueDeclare) {
	log.Printf("value is: %+v\n", queue)
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
	log.Printf("value is: %+v\n", exchange)
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
	log.Printf("value is: %+v\n", bind)
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
	log.Printf("value is: %+v\n", publish)
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
	log.Printf("value is: %+v\n", consume)
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
				time.Sleep(1 * time.Millisecond) //slows down loop
				select {
					case msg := <-msgs:
						body := string(msg.Body[:])
						arr = append(arr,body)
						end = false
					default:
						end = true
				}
				if(end){
					break
				}
			}
			log.Printf("messages len=%d cap=%d %v\n", len(arr), cap(arr), arr)
			json, readError := json.Marshal(arr)
			if !checkError(w,readError){
				w.Write([]byte(string(json)))
			}else{
				log.Println("error converting to json\n")
			}
		}
	}
}


func RabbitSubscribe(w http.ResponseWriter, subscribe models.QueueSubscribe) {
	ch, conn := createChanel(w)
	if ch != nil {
		qos := subscribe.Qos
		err := ch.Qos(
			qos.PrefetchCount,
			qos.PrefetchSize,
			false,
		)
		if !checkError(w,err){
			msgs, err := ch.Consume(
				subscribe.Name,
				subscribe.Consumer,
				false,
				subscribe.Exclusive,
				subscribe.NoLocal,
				subscribe.NoWait,
				subscribe.Arguments,
			)
			if !checkError(w,err){
				subscribeInit(w, msgs, subscribe.URL, subscribe.Timeout, ch, conn)
				return
			}
		}
		close(ch,conn)
	}
}

func subscribeInit(w http.ResponseWriter, msgs <-chan amqp.Delivery, url string,
	timeout time.Duration, ch *amqp.Channel, conn *amqp.Connection) {
		id,err := util.NewUUID()
		log.Println(id)
		if !checkError(w,err){
			consumers[id] = make(chan bool,1)
			for i:= 0 ; i < configs.Threads ; i++ {
				go subscribeLoop(msgs, id, url, timeout, ch, conn)
			}
			subscribe := models.QueueSubscribeId {
				ID: id,
			}
			json, readError := json.Marshal(&subscribe)
			if !checkError(w,readError){
				log.Println("new subscriber: ",string(json))
				w.Write([]byte(string(json)))
				return
			}
		}
		close(ch,conn)
}

func subscribeLoop(msgs <-chan amqp.Delivery, id string, url string,
	timeout time.Duration, ch *amqp.Channel, conn *amqp.Connection){
	b := consumers[id]
	for {
		time.Sleep(1 * time.Second) //slows down loop
		select {
			case v,_ := <- b:
				delete(consumers,id) //clear reference outside thread
				b <- v               //stops other threads
				close(ch, conn)
				log.Println("closed connection ",id)
				return
			case msg := <-msgs:
				go send(id, msg, url, timeout)
			default:
				continue
		}
	}
}

func RabbitUnsubscribe(id string){
	if b := consumers[id] ; b != nil {
		consumers[id] <- false
	}
}

type data struct {
	Consumers               map[string](chan bool)     `json:"consumers"`
	FailedMessages          map[uint64](uint64)        `json:"failedmessages"`
}

func ArrayDump(w http.ResponseWriter, password string){
	pass := os.Getenv("DUMP_PASS")
	if pass != "" && password == pass {
		dump := data{
			Consumers: consumers,
			FailedMessages: messages,
		}
		json, err := json.Marshal(&dump)
		if !checkError(w,err){
			log.Println("dump: ",string(json))
			w.Write([]byte(string(json)))
		}
	}else{
		w.WriteHeader(403)
	}
}

func send(id string, msg amqp.Delivery, url string, timeout time.Duration){
	m := models.QueueMessage{
		ID: msg.DeliveryTag,
		Body: msg.Body,
	}
	message,err := json.Marshal(&m)
	if err != nil {
		log.Printf("failed to make json [%s]",m)
	}
	client := http.Client{
		Timeout: timeout,
	}
	buffer := bytes.NewBuffer(message)
	resp,err1 := client.Post(url,"application/json",buffer)
	defer resp.Body.Close()
	if err1 != nil {
		log.Println(err.Error())
	}else if resp != nil {
		switch resp.StatusCode {
			case 200:
				return
			case 404:
				RabbitUnsubscribe(id)
				log.Printf("Error: [404] failed to post ending subscription id [%s]",id)
			default:
				messages[msg.DeliveryTag] += 1
				log.Printf("failed to post message %d %v",resp.StatusCode,resp)
		}
	}
}

func close(ch *amqp.Channel, conn *amqp.Connection){
	conn.Close()
	ch.Close()
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