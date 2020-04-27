package dep

import (
	"github.com/ProjectReferral/Get-me-in/customer-api/configs"
	"github.com/ProjectReferral/Get-me-in/customer-api/internal/api"
	"github.com/ProjectReferral/Get-me-in/customer-api/internal/event-driven"
	"github.com/ProjectReferral/Get-me-in/customer-api/lib/hermes"
	"github.com/ProjectReferral/Get-me-in/queueing-api/client"
	"github.com/ProjectReferral/Get-me-in/queueing-api/client/models"
	"github.com/gorilla/mux"
	"log"
	"time"
)

//methods that are implemented on util
//and will be used
type ConfigBuilder interface {
	LoadEnvConfigs()
}

//internal specific configs are loaded at runtime
//takes in a object(implemented interface) of type ServiceConfigs
func Inject(builder ConfigBuilder) {

	//load the env into the object
	builder.LoadEnvConfigs()

	//dependency injection to our resource
	//we inject the rabbitmq client
	rabbitMQClient := builder.LoadRabbitMQConfigs()
	loadRabbitMQClient(rabbitMQClient)

	subscribeToChannels()

	eb := api.EndpointBuilder{}

	eb.InjectService(&hermes.EmailStruct{})
	eb.SetupRouter(mux.NewRouter())
	eb.SetupEndpoints()

	eb.SetupSubscriptionEndpoint(rabbitMQClient)
}

func loadRabbitMQClient(q client.QueueClient) {
	log.Println("Injecting RabbitMQ Repo")
	event_driven.Client = q
}

func subscribeToChannels() {
	log.Println("Subscribing to channels ...")
	event_driven.SubscribeTo(models.QueueSubscribe{
		//endpoint that will be consuming the messages
		URL:       configs.SUB_ACTION_EMAIL,
		Name:      "new-user-verify-email",
		Consumer:  "",
		Exclusive: false,
		NoLocal:   false,
		NoWait:    false,
		MaxRetry:  0,
		Timeout:   5 * time.Second,
		Qos: models.QueueQos{
			PrefetchCount: 0,
			PrefetchSize:  0,
		},
	})
}
