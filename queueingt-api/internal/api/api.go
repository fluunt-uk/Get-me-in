package api

import (
	"encoding/json"
	"net/http"
	events "github.com/ProjectReferral/Get-me-in/queueingt-api/internal/event-driven"
	"github.com/ProjectReferral/Get-me-in/queueingt-api/internal/models"
)

func TestFunc(w http.ResponseWriter, r *http.Request) {
	if events.TestQ() {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func CreateQueue(w http.ResponseWriter, r *http.Request) {
	if validateBody(w, r) {
		queue := models.QueueDeclare{Arguments: nil}
		send(w, r, queue, events.RabbitCreateQueue)
	}
}

func CreateExchange(w http.ResponseWriter, r *http.Request) {
	if validateBody(w, r) {
		exchange := models.ExchangeDeclare{Arguments: nil}
		send(w, r, exchange, events.RabbitCreateExchange)
	}
}

func BindExchange(w http.ResponseWriter, r *http.Request) {
	if validateBody(w, r) {
		bind := models.QueueBind{Arguments: nil}
		send(w, r, bind, events.RabbitQueueBind)
	}
}

func PublishToExchange(w http.ResponseWriter, r *http.Request) {
	if validateBody(w, r) {
		publish := models.ExchangePublish{}
		send(w, r, publish, events.RabbitPublish)
	}
}


func validateBody(w http.ResponseWriter, r *http.Request) bool {
	if r.ContentLength < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No body error!"))
		return false
	}
	return true
}

func send(w http.ResponseWriter, r *http.Request,
	s interface{}, fn func(interface{})) {
	err := json.NewDecoder(r.Body).Decode(&s)
	if !HandleError(err, w) {
		fn(s)
	}
}
