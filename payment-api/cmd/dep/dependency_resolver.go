package dep

import (
	"github.com/ProjectReferral/Get-me-in/payment-api/lib/dynamodb/repo"
	"github.com/ProjectReferral/Get-me-in/payment-api/lib/rabbitmq"
	"github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api/resources/card"
	customer "github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api/resources/customer"
	sub "github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api/resources/subscription"
	token "github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api/resources/token"
	"github.com/ProjectReferral/Get-me-in/pkg/dynamodb"
	"github.com/ProjectReferral/Get-me-in/queueing-api/client"
	"log"
)

//methods that are implemented on util
//and will be used
type ConfigBuilder interface{
	LoadEnvConfigs()
	LoadDynamoDBConfigs() *dynamodb.Wrapper
	LoadRabbitMQConfigs() *client.DefaultQueueClient
}

func Inject(builder ConfigBuilder){

	builder.LoadEnvConfigs()

	//setup dynamo library
	dynamoClient := builder.LoadDynamoDBConfigs()
	//connect to the instance
	log.Println("Connecting to Dynamo Client")
	dynamoClient.DefaultConnect()

	rabbitMQClient := builder.LoadRabbitMQConfigs()
	//dependency injection to our resource
	//we inject the rabbitmq client
	LoadRabbitMQClient(rabbitMQClient)

	LoadCardClient(&card.Wrapper{})
	LoadCustomerClient(&customer.Wrapper{})
	LoadSubClient(&sub.Wrapper{DynamoSubRepo:&repo.Wrapper{DC:dynamoClient}})
	LoadTokenClient(&token.Wrapper{})
}

func LoadRabbitMQClient(c client.QueueClient){
	log.Println("Injecting RabbitMQ Client")
	rabbitmq.Client = c
}