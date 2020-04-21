package dep

import (
	"github.com/ProjectReferral/Get-me-in/marketing-api/lib/dynamodb/repo-builder"
	"github.com/ProjectReferral/Get-me-in/pkg/dynamodb"
	"log"
)

//methods that are implemented on util
//and will be used
type ConfigBuilder interface{
	LoadEnvConfigs()
	LoadDynamoDBConfigs() *dynamodb.Wrapper
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

}

//variable injected with the interface methods
func LoadSignInRepo (r repo_builder.SignInBuilder){
	log.Println("Injecting SignIn Repo")
	repo_builder.SignIn = r
}

