package dep

import (
	sub_builder "github.com/ProjectReferral/Get-me-in/payment-api/lib/dynamodb/sub-builder"
	"github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api/resources/card"
	"github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api/resources/subscription"
	"github.com/ProjectReferral/Get-me-in/pkg/dynamodb"
	"log"
)

//methods that are implemented on util
//and will be used
type ConfigBuilder interface{
	LoadEnvConfigs()
	LoadDynamoDBConfigs() *dynamodb.Wrapper
}

func Inject(builder ConfigBuilder){

	builder.LoadEnvConfigs()

	//setup dynamo library
	//TODO:shall the dynamo configs injected here? or in the main?
	dynamoClient := builder.LoadDynamoDBConfigs()
	//connect to the instance
	log.Println("Connecting to Dynamo Client")
	dynamoClient.DefaultConnect()

	//dependency injection to our resource
	//we inject the dynamo client
	//shared client, therefore shared in between all the repos
	LoadSubRepo(&sub_builder.Wrapper{})

	LoadCardService(&card.Wrapper{})
}


//variable injected with the interface methods
func LoadSubRepo (r sub_builder.Builder){
	log.Println("Injecting Sub Repo")
	subscription.SubRepo = r
}
