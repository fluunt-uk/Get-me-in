package dep

import (
	event_driven "github.com/ProjectReferral/Get-me-in/account-api/internal/event-driven"
	"github.com/ProjectReferral/Get-me-in/account-api/lib/dynamodb/repo"
	"github.com/ProjectReferral/Get-me-in/pkg/dynamodb"
	"log"
	"os"
)

//methods that are implemented on util
//and will be used
type Loader interface{
	LoadEnvConfigs()
	LoadDynamoDBConfigs() *dynamodb.DynamoDB
}

//variable injected with the interface methods
func LoadAccountRepo (r repo.AccountRepository){
	log.Println("Injecting Account repo")
	repo.Account = r
}
//variable injected with the interface methods
func LoadAccountAdvertRepo (r repo.AccountAdvertRepository){
	log.Println("Injecting Account Advert Repo")
	repo.AccountAdvert = r
}
//variable injected with the interface methods
func LoadSignInRepo (r repo.SignInRepository){
	log.Println("Injecting SignIn Repo")
	repo.SignIn = r
}

//internal specific configs are loaded at runtime
//takes in a object(implemented interface) of type ServiceConfigs
func Inject(loader Loader) {

	//load the env into the object
	loader.LoadEnvConfigs()

	//setup dynamo library
	//TODO:shall the dynamo configs injected here? or in the main?
	dynamoClient := loader.LoadDynamoDBConfigs()
	//connect to the instance
	log.Println("Connecting to dynamo client")
	dynamoClient.Connect()

	//dependency injection to our resource
	//we inject the dynamo client
	//shared client, therefore shared in between all the repos
	LoadSignInRepo(&repo.DynamoSignIn{
		DC: dynamoClient,
	})

	LoadAccountRepo(&repo.DynamoAccount{
		DC: dynamoClient,
	})

	LoadAccountAdvertRepo(&repo.DynamoAccountAdvert{
		DC: dynamoClient,
	})

	//dependency injection to our resource
	//we inject the rabbitmq client
	//TODO: will be done through the network(REST API)
	event_driven.MQ = &event_driven.RabbitClient{
		URL:        os.Getenv("BROKERURL"),
	}

}
////TODO: can be used for unit testing?
//type Repository interface {
//	AccountRepository
//	AccountAdvertRepository
//	SignInRepository
//}