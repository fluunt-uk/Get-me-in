package event_driven

import (
	"github.com/ProjectReferral/Get-me-in/queueingt-api/internal/models"
	"fmt"
)

func RabbitCreateQueue(i interface{}) {
	queue,ok := i.(models.QueueDeclare)
	if ok {
		fmt.Printf("value is: %+v\n", queue)
	}
}

func RabbitCreateExchange(i interface{}) {
	queue,ok := i.(models.ExchangeDeclare)
	if ok {
		fmt.Printf("value is: %+v\n", queue)
	}
}

func RabbitQueueBind(i interface{}) {
	queue,ok := i.(models.QueueBind)
	if ok {
		fmt.Printf("value is: %+v\n", queue)
	}
}

func RabbitPublish(i interface{}) {
	queue,ok := i.(models.ExchangePublish)
	if ok {
		fmt.Printf("value is: %+v\n", queue)
	}
}
