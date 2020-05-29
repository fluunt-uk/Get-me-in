package dep

import (
	"github.com/ProjectReferral/Get-me-in/customer-api/configs"
	"github.com/ProjectReferral/Get-me-in/customer-api/internal/api"
	ed "github.com/ProjectReferral/Get-me-in/customer-api/internal/event-driven"
	"github.com/ProjectReferral/Get-me-in/customer-api/internal/service"
	"github.com/ProjectReferral/Get-me-in/customer-api/lib/hermes/templates"
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
	LoadRabbitMQConfigs() *client.DefaultQueueClient
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

	log.Println("Setting up message handler...")
	//initialise our message handler
	mh := &ed.MsgHandler{}

	es := &service.EmailService{EB: &templates.EmailBuilder{}}
	es.SetupTemplates()
	//inject the hermes service into it
	mh.InjectService(es)

	log.Println("Loading endpoints...")
	eb := api.EndpointBuilder{}

	eb.SetupRouter(mux.NewRouter())
	eb.SetupEndpoints()

	eb.SetQueueClient(rabbitMQClient)
	// we use the message handler here
	eb.SetupMsgHandler(mh)
	eb.SetupSubscriptionEndpoint()
}

func loadRabbitMQClient(q client.QueueClient) {
	log.Println("Injecting RabbitMQ Client and dependencies")
	ed.Client = q
	ss := &ed.SubscriberStore{}
	ss.Init()
	ed.Store = ss
}

func subscribeToChannels() {
	log.Println("Subscribing to channels ...")

	//subscribe to new user verification email Q
	ed.SubscribeTo(models.QueueSubscribe{
		//endpoint that will be consuming the messages
		URL:       configs.EMAIL_DISPATCHER_URL,
		Name:      configs.VERIFY_EMAIL_Q,
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
