package event_driven

import (
	"github.com/ProjectReferral/Get-me-in/queueing-api/configs"
	"github.com/ProjectReferral/Get-me-in/queueing-api/client/models"
	"github.com/ProjectReferral/Get-me-in/util"
	"github.com/streadway/amqp"
	"log"
	"encoding/json"
	"strings"
	"strconv"
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

func CreateFailedMessageQueue(){
	args := amqp.Table{}
	args["x-queue-type"] = "classic"
	queue := models.QueueDeclare{
		Name: configs.FailedMessageQueue,
		Durable: true,
		DeleteWhenUnused: false,
		Exclusive: false,
		NoWait: false,
		Arguments: args,
	}
	RabbitCreateQueue(nil,queue,true)
}

func RabbitCreateQueue(w http.ResponseWriter, queue models.QueueDeclare, fatal bool) {
	log.Printf("value is: %+v\n", queue)
	ch, conn := createChanel(w,fatal)
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
		checkError(w,err,fatal)
	}
}

func RabbitCreateExchange(w http.ResponseWriter, exchange models.ExchangeDeclare) {
	log.Printf("value is: %+v\n", exchange)
	ch, conn := createChanel(w,false)
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
		checkError(w,err,false)
	}
}

func RabbitQueueBind(w http.ResponseWriter, bind models.QueueBind) {
	log.Printf("value is: %+v\n", bind)
	ch, conn := createChanel(w,false)
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
		checkError(w,err,false)
	}
}

func RabbitPublish(w http.ResponseWriter, publish models.ExchangePublish) {
	log.Printf("value is: %+v\n", publish)
	ch, conn := createChanel(w,false)
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
		checkError(w,err,false)
	}
}

func RabbitConsume(w http.ResponseWriter, consume models.QueueConsume) {
	log.Printf("value is: %+v\n", consume)
	ch, conn := createChanel(w,false)
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
		if !checkError(w,err,false){
			var arr = []string{}
			var end bool
			for {
				time.Sleep(configs.SleepTime) //slows down loop
				select {
					case msg := <-msgs:
						log.Printf("value %+v",msg)
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
			w.Header().Set("Content-Type", "application/json")
			jsonErr := json.NewEncoder(w).Encode(arr)
			checkError(w,jsonErr,false)
		}
	}
}


func RabbitSubscribe(w http.ResponseWriter, subscribe models.QueueSubscribe) {
	ch, conn := createChanel(w,false)
	if ch != nil {
		qos := subscribe.Qos
		err := ch.Qos(
			qos.PrefetchCount,
			qos.PrefetchSize,
			false,
		)
		if !checkError(w,err,false){
			msgs, err := ch.Consume(
				subscribe.Name,
				subscribe.Consumer,
				false,
				subscribe.Exclusive,
				subscribe.NoLocal,
				subscribe.NoWait,
				subscribe.Arguments,
			)
			if !checkError(w,err,false){
				subscribeInit(w, msgs, subscribe.URL, subscribe.MaxRetry, subscribe.Timeout, ch, conn)
				return
			}
		}
		close(ch,conn)
	}
}

func RabbitUnsubscribe(id string){
	if b := consumers[id] ; b != nil {
		consumers[id] <- false
	}
}

func RabbitAck(w http.ResponseWriter, ma models.MessageAcknowledge){
	log.Printf("value is: %+v\n", ma)
	ch, conn := createChanel(w,false)
	if ch != nil {
		defer conn.Close()
		defer ch.Close()
		var err error
		if ma.Acknowledge {
			err = ch.Ack(
				ma.ID,
				ma.Multiple,
			)
		}else {
			err = ch.Nack(
				ma.ID,
				ma.Multiple,
				ma.Requeue,
			)
		}
		if !ma.Acknowledge && ma.Requeue {
			messages[ma.ID] += 1
		}
		checkError(w,err,false)
	}
}

func RabbitReject(w http.ResponseWriter, mr models.MessageReject){
	log.Printf("value is: %+v\n", mr)
	ch, conn := createChanel(w,false)
	if ch != nil {
		defer conn.Close()
		defer ch.Close()
		err := ch.Ack(
				mr.ID,
				mr.Requeue,
		)
		checkError(w,err,false)
		if mr.Requeue {
			messages[mr.ID] += 1
		}
	}
}

func subscribeInit(w http.ResponseWriter, msgs <-chan amqp.Delivery, url string,
	maxRetry int, timeout time.Duration, ch *amqp.Channel, conn *amqp.Connection) {
		id,err := util.NewUUID()
		log.Println(id)
		if !checkError(w,err,false){
			consumers[id] = make(chan bool,1)
			for i:= 0 ; i < configs.Threads ; i++ {
				go subscribeLoop(msgs, id, url, maxRetry, timeout, ch, conn)
			}
			subscribe := models.QueueSubscribeId {
				ID: id,
			}
			w.Header().Set("Content-Type", "application/json")
			jsonErr := json.NewEncoder(w).Encode(subscribe)
			checkError(w,jsonErr,false)
			log.Println("new subscriber: %s",id)
			return
		}
		close(ch,conn)
}

func subscribeLoop(msgs <-chan amqp.Delivery, id string, url string, maxRetry int,
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
				if !maxRetryExceeded(msg, messages[msg.DeliveryTag], maxRetry, ch) {
					go send(id, msg, url, timeout)
				}
			default:
				continue
		}
	}
}

func maxRetryExceeded(msg amqp.Delivery, count uint64, max int, ch *amqp.Channel) bool {
	if count > uint64(max) {
		return rejectMessage(msg, messages[msg.DeliveryTag], max, ch)
	}
	return false
}

func rejectMessage(msg amqp.Delivery, count uint64, max int, ch *amqp.Channel) bool {
	id := msg.DeliveryTag
	body := msg.Body
	log.Printf("rejected message id [%d]",id)
	m := models.QueueFailedMessage{
		ID: id,
		Body: body,
		RetryCount: count,
		Reason: "Message requeued more than " + strconv.Itoa(max),
	}
	json, jsonError := json.Marshal(&m)
	if !checkError(nil,jsonError,false){
		publish := amqp.Publishing{
			Body: []byte(json),
		}
		err := ch.Publish(
			configs.FailedMessageQueue,
			"",
			false,
			false,
			publish,
		)
		if !checkError(nil,err,false){
			if err1 := msg.Reject(false) ; err1 != nil {
				return true
			}
		}
	}
	return false
}

func send(id string, msg amqp.Delivery, url string, timeout time.Duration){
	m := models.QueueMessage{
		ID: msg.DeliveryTag,
		Body: msg.Body,
	}
	message,err := json.Marshal(&m)
	if err != nil {
		log.Printf("failed to make json [%+v]",m)
	}
	client := &http.Client{
		Timeout: timeout,
	}
	buffer := bytes.NewBuffer(message)
	if(consumers[id] == nil){
		return // last minute unsubscribe check
	}
	resp,errP := client.Post(url,"application/json",buffer)
	if errP != nil {
		log.Println(errP)
		return
	}
	if resp != nil {
		defer resp.Body.Close()
		switch resp.StatusCode {
			case 200:
				return
			case 404:
				RabbitUnsubscribe(id)
				log.Printf("Error: [404] ending subscription id [%s]",id)
			default:
				messages[msg.DeliveryTag] += 1
				log.Printf("failed to post message %d %v",resp.StatusCode,resp)
		}
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
		log.Println("dump: %+v",dump)
		w.Header().Set("Content-Type", "application/json")
		jsonErr := json.NewEncoder(w).Encode(dump)
		checkError(w,jsonErr,false)
	}else{
		w.WriteHeader(403)
	}
}

func close(ch *amqp.Channel, conn *amqp.Connection){
	conn.Close()
	ch.Close()
}

func createChanel(w http.ResponseWriter, fatal bool) (*amqp.Channel, *amqp.Connection) {
	conn := createConnection(w,fatal)
	if conn != nil {
		ch, err := conn.Channel()
		if !checkError(w,err,fatal) {
			return ch,conn
		}else{
			defer conn.Close()
		}
	}
	return nil,nil
}

func createConnection(w http.ResponseWriter, fatal bool) *amqp.Connection {
	conn, err := amqp.Dial(configs.BrokerUrl)
	if checkError(w,err,fatal) {
		return nil
	}
	return conn
}

func checkError(w http.ResponseWriter, err error, fatal bool) bool{
	if err != nil {
		log.Println(err.Error())
		status := 400
		if strings.Contains(err.Error(),"Exception (404)") {
			status = 404
		}else{
			status = 400
			
		}
		if w != nil {
			w.WriteHeader(status)
		}else if fatal {
			log.Fatalf("%s: %+v", "Failed to init service", err)
		}
		return true
	}
	return false
}