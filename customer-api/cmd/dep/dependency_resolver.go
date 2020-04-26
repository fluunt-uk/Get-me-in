package dep

import (
	"github.com/ProjectReferral/Get-me-in/customer-api/internal/api"
	"github.com/ProjectReferral/Get-me-in/customer-api/internal/api/email"
	email_builder "github.com/ProjectReferral/Get-me-in/customer-api/internal/event-driven"
	"github.com/ProjectReferral/Get-me-in/queueing-api/client"
	"log"
	"net/http"
	"time"
)

//methods that are implemented on util
//and will be used
type ConfigBuilder interface{
	LoadEnvConfigs()
	LoadRabbitMQConfigs() *client.DefaultQueueClient
}

//internal specific configs are loaded at runtime
//takes in a object(implemented interface) of type ServiceConfigs
func Inject(builder ConfigBuilder) {

	//load the env into the object
	builder.LoadEnvConfigs()


	hc := &http.Client{Timeout: 5 * time.Second}

	//dependency injection to our resource
	//we inject the rabbitmq client
	rabbitMQClient := builder.LoadRabbitMQConfigs()
	email_builder.Subscribe(rabbitMQClient, hc)

	email_builder.SetupRoute(&api.Router, hc, rabbitMQClient)

	LoadEmailRepo(&email_builder.EmailStruct{})
}
//variable injected with the interface methods
func LoadEmailRepo (r email_builder.EmailBuilder){
	log.Println("Injecting Email Repo")
	email.Service = r
}