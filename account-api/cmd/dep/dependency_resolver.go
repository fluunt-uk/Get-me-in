package dep

import (
	"github.com/ProjectReferral/Get-me-in/account-api/lib/dynamodb/repo-builder"
	"github.com/ProjectReferral/Get-me-in/account-api/lib/rabbitmq"
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

//internal specific configs are loaded at runtime
//takes in a object(implemented interface) of type ServiceConfigs
func Inject(builder ConfigBuilder) {

	//load the env into the object
	builder.LoadEnvConfigs()

	//setup dynamo library
	//TODO:shall the dynamo configs injected here? or in the main?
	dynamoClient := builder.LoadDynamoDBConfigs()
	//connect to the instance
	log.Println("Connecting to dynamo client")
	dynamoClient.DefaultConnect()

	//dependency injection to our resource
	//we inject the dynamo client
	//shared client, therefore shared in between all the repos
	LoadSignInRepo(&repo_builder.SignInWrapper{
		DC: dynamoClient,
	})

	LoadAccountRepo(&repo_builder.AccountWrapper{
		DC: dynamoClient,
	})

	LoadAccountAdvertRepo(&repo_builder.AccountAdvertWrapper{
		DC: dynamoClient,
	})

	//dependency injection to our resource
	//we inject the rabbitmq client
	rabbitMQClient := builder.LoadRabbitMQConfigs()

	LoadRabbitMQClient(rabbitMQClient)
}

//variable injected with the interface methods
func LoadAccountRepo (r repo_builder.AccountBuilder){
	log.Println("Injecting Account repo")
	repo_builder.Account = r
}
//variable injected with the interface methods
func LoadAccountAdvertRepo (r repo_builder.AccountAdvertBuilder){
	log.Println("Injecting Account Advert Repo")
	repo_builder.AccountAdvert = r
}
//variable injected with the interface methods
func LoadSignInRepo (r repo_builder.SignInBuilder){
	log.Println("Injecting SignIn Repo")
	repo_builder.SignIn = r
}

func LoadRabbitMQClient(c client.QueueClient){
	log.Println("Injecting RabbitMQ Client")
	rabbitmq.Client = c
}

