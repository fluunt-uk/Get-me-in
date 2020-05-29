package client

import (
	"bytes"
	"encoding/json"
	"github.com/ProjectReferral/Get-me-in/queueing-api/client/models"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type HttpReponse func(*http.Response, error)

type HandleMessage func(qm *models.QueueMessage, err error, qc QueueClient) (sm models.SubscribeMessage, hr HttpReponse)

type QueueClient interface {
	SetupURL(url string)

	GetURL() string

	CreateQueue(client *http.Client, queue models.QueueDeclare) (resp *http.Response, err error)

	CreateExchange(client *http.Client, exchange models.ExchangeDeclare) (resp *http.Response, err error)

	QueueBind(client *http.Client, bind models.QueueBind) (resp *http.Response, err error)

	Publish(client *http.Client, publish models.ExchangePublish) (resp *http.Response, err error)

	Consume(client *http.Client, consume models.QueueConsume) (resp *http.Response, err error)

	SetupRoute(router *mux.Router, route string, client *http.Client, hm HandleMessage)

	Subscribe(client *http.Client, subscribe models.QueueSubscribe) (resp *http.Response, err error)

	UnSubscribe(client *http.Client, subscribeID models.QueueSubscribeId) (resp *http.Response, err error)

	acknowledge(client *http.Client, acknowledge models.MessageAcknowledge) (resp *http.Response, err error)

	reject(client *http.Client, reject models.MessageReject) (resp *http.Response, err error)
}

type DefaultQueueClient struct {
	url string
}

//pointer as it is a setter
func (dqc *DefaultQueueClient) SetupURL(url string) {
	dqc.url = url
}

func (dqc *DefaultQueueClient) GetURL() string {
	return dqc.url
}

func (dqc *DefaultQueueClient) CreateQueue(client *http.Client, queue models.QueueDeclare) (resp *http.Response, err error) {
	body, err := json.Marshal(queue)
	if err != nil {
		log.Printf("failed to make json [%v]", queue)
		return nil, err
	}
	return client.Post(dqc.url+"/queue", "application/json", bytes.NewBuffer(body))
}

func (dqc *DefaultQueueClient) CreateExchange(client *http.Client, exchange models.ExchangeDeclare) (resp *http.Response, err error) {
	body, err := json.Marshal(exchange)
	if err != nil {
		log.Printf("failed to make json [%v]", exchange)
		return nil, err
	}
	return client.Post(dqc.url+"/exchange", "application/json", bytes.NewBuffer(body))
}

func (dqc *DefaultQueueClient) QueueBind(client *http.Client, bind models.QueueBind) (resp *http.Response, err error) {
	body, err := json.Marshal(bind)
	if err != nil {
		log.Printf("failed to make json [%v]", bind)
		return nil, err
	}
	PutReq, err1 := http.NewRequest("PUT", dqc.url+"/bind", bytes.NewBuffer(body))
	if err1 != nil {
		log.Printf("failed to create request: %s", err1.Error())
		return nil, err1
	}
	PutReq.Header.Set("Content-Type", "application/json")
	return client.Do(PutReq)
}

func (dqc *DefaultQueueClient) Publish(client *http.Client, publish models.ExchangePublish) (resp *http.Response, err error) {
	body, err := json.Marshal(publish)
	if err != nil {
		log.Printf("failed to make json [%v]", publish)
		return nil, err
	}
	return client.Post(dqc.url+"/publish", "application/json", bytes.NewBuffer(body))
}

func (dqc *DefaultQueueClient) Consume(client *http.Client, consume models.QueueConsume) (resp *http.Response, err error) {
	body, err := json.Marshal(consume)
	if err != nil {
		log.Printf("failed to make json [%v]", consume)
		return nil, err
	}
	return client.Post(dqc.url+"/consume", "application/json", bytes.NewBuffer(body))
}

func (dqc *DefaultQueueClient) SetupRoute(router *mux.Router, route string, client *http.Client, hm HandleMessage) {
	router.HandleFunc(route, handleResponse(hm, client, dqc)).Methods("POST")
}

func (dqc DefaultQueueClient) Subscribe(client *http.Client, subscribe models.QueueSubscribe) (resp *http.Response, err error) {
	body, err := json.Marshal(subscribe)
	if err != nil {
		log.Printf("failed to make json [%v]", subscribe)
		return nil, err
	}
	return client.Post(dqc.url+"/subscribe", "application/json", bytes.NewBuffer(body))
}

func (dqc *DefaultQueueClient) UnSubscribe(client *http.Client, subscribeID models.QueueSubscribeId) (resp *http.Response, err error) {
	body, err := json.Marshal(subscribeID)
	if err != nil {
		log.Printf("failed to make json [%v]", subscribeID)
		return nil, err
	}
	return client.Post(dqc.url+"/unsubscribe", "application/json", bytes.NewBuffer(body))
}

func (dqc *DefaultQueueClient) acknowledge(client *http.Client, acknowledge models.MessageAcknowledge) (resp *http.Response, err error) {
	body, err := json.Marshal(acknowledge)
	if err != nil {
		log.Printf("failed to make json [%v]", acknowledge)
		return nil, err
	}
	return client.Post(dqc.url+"/acknowledge", "application/json", bytes.NewBuffer(body))
}

func (dqc *DefaultQueueClient) reject(client *http.Client, reject models.MessageReject) (resp *http.Response, err error) {
	body, err := json.Marshal(reject)
	if err != nil {
		log.Printf("failed to make json [%v]", reject)
		return nil, err
	}
	return client.Post(dqc.url+"/reject", "application/json", bytes.NewBuffer(body))
}

func ExtractJsonString(r *http.Request) (json string, err error) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("failed to make json string [%v]", err)
		return string(data), nil
	}
	return "", err
}

func handleResponse(handleMessage HandleMessage, client *http.Client, qc QueueClient) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		message := models.QueueMessage{}
		jsonErr := json.NewDecoder(req.Body).Decode(&message)
		if req != nil {
			defer req.Body.Close()
		}
		if jsonErr != nil {
			w.Write([]byte(jsonErr.Error()))
			w.WriteHeader(400)
		}
		go func() {
			result, handler := handleMessage(&message, jsonErr, qc) //run in another thread to avoid blocking
			if result != nil && handler != nil {
				if ma, ok := result.(models.MessageAcknowledge); ok {
					ma.Body = message.Body
					handler(qc.acknowledge(client, ma))
				}
				if mr, ok := result.(models.MessageReject); ok {
					mr.Body = message.Body
					handler(qc.reject(client, mr))
				}
			}
		}()
	}
}
