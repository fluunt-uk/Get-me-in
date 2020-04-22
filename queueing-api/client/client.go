package client

import (
	"github.com/ProjectReferral/Get-me-in/queueing-api/client/models"
	"bytes"
	"log"
	"encoding/json"
	"net/http"
	"io/ioutil"
)

type QueueClient interface {
	SetupURL(url string)
	
	CreateQueue(client *http.Client, queue models.QueueDeclare) (resp *http.Response, err error)
	
	CreateExchange(client *http.Client, exchange models.ExchangeDeclare) (resp *http.Response, err error)
	
	QueueBind(client *http.Client, bind models.QueueBind) (resp *http.Response, err error)
	
	Publish(client *http.Client, publish models.ExchangePublish) (resp *http.Response, err error)
	
	Consume(client *http.Client, consume models.QueueConsume) (resp *http.Response, err error)
}

type DefaultQueueClient struct {
	url string
}

//pointer as it is a setter
func (dqc *DefaultQueueClient) SetupURL(url string){
	dqc.url = url
}

func (dqc DefaultQueueClient) CreateQueue(client *http.Client, queue models.QueueDeclare) (resp *http.Response, err error) {
	body,err := json.Marshal(queue)
	if err != nil {
		log.Printf("failed to make json [%v]",queue)
		return nil,err
	}
	return client.Post(dqc.url + "/queue","application/json",bytes.NewBuffer(body))
}

func (dqc DefaultQueueClient) CreateExchange(client *http.Client, exchange models.ExchangeDeclare) (resp *http.Response, err error) {
	body,err := json.Marshal(exchange)
	if err != nil {
		log.Printf("failed to make json [%v]",exchange)
		return nil,err
	}
	return client.Post(dqc.url + "/exchange","application/json",bytes.NewBuffer(body))
}

func (dqc DefaultQueueClient) QueueBind(client *http.Client, bind models.QueueBind) (resp *http.Response, err error) {
	body,err := json.Marshal(bind)
	if err != nil {
		log.Printf("failed to make json [%v]",bind)
		return nil,err
	}
    PutReq, err1 := http.NewRequest("PUT", dqc.url + "/bind", bytes.NewBuffer(body))
	if err1 != nil {
		log.Printf("failed to create request: %s",err1.Error())
		return nil,err1
	}
    PutReq.Header.Set("Content-Type", "application/json")
    return client.Do(PutReq)
}

func (dqc DefaultQueueClient) Publish(client *http.Client, publish models.ExchangePublish) (resp *http.Response, err error) {
	body,err := json.Marshal(publish)
	if err != nil {
		log.Printf("failed to make json [%v]",publish)
		return nil,err
	}
	return client.Post(dqc.url + "/publish","application/json",bytes.NewBuffer(body))
}

func (dqc DefaultQueueClient) Consume(client *http.Client, consume models.QueueConsume) (resp *http.Response, err error) {
	body,err := json.Marshal(consume)
	if err != nil {
		log.Printf("failed to make json [%v]",consume)
		return nil,err
	}
	return client.Post(dqc.url + "/consume","application/json",bytes.NewBuffer(body))
}

func ExtractJsonString(r *http.Request) (json string, err error) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil{
		log.Printf("failed to make json string [%v]",err)
		return string(data), nil
	}
	return "", err
}