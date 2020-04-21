package api

import (
	"encoding/json"
	"net/http"
	events "github.com/ProjectReferral/Get-me-in/queueing-api/internal/event-driven"
	"github.com/ProjectReferral/Get-me-in/queueing-api/client/models"
)

func TestFunc(w http.ResponseWriter, r *http.Request) {
	if events.TestQ(w) {
		w.WriteHeader(http.StatusOK)
	}
}

func CreateQueue(w http.ResponseWriter, r *http.Request) {
	queue := models.QueueDeclare{Arguments: nil}
	err := json.NewDecoder(r.Body).Decode(&queue)
	if !HandleError(err, w) {
		events.RabbitCreateQueue(w,queue)
	}
}

func CreateExchange(w http.ResponseWriter, r *http.Request) {
	exchange := models.ExchangeDeclare{Arguments: nil}
	err := json.NewDecoder(r.Body).Decode(&exchange)
	if !HandleError(err, w) {
		events.RabbitCreateExchange(w,exchange)
	}
}

func BindExchange(w http.ResponseWriter, r *http.Request) {
	bind := models.QueueBind{Arguments: nil}
	err := json.NewDecoder(r.Body).Decode(&bind)
	if !HandleError(err, w) {
		events.RabbitQueueBind(w,bind)
	}
}

func PublishToExchange(w http.ResponseWriter, r *http.Request) {
	publish := models.ExchangePublish{}
	err := json.NewDecoder(r.Body).Decode(&publish)
	if !HandleError(err, w) {
		events.RabbitPublish(w,publish)
	}
}

func ConsumeQueue(w http.ResponseWriter, r *http.Request) {
	consume := models.QueueConsume{Arguments: nil}
	err := json.NewDecoder(r.Body).Decode(&consume)
	if !HandleError(err, w) {
		events.RabbitConsume(w,consume)
	}
}