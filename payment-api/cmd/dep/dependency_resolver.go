package dep

import (
	"github.com/ProjectReferral/Get-me-in/payment-api/lib/dynamodb/repo"
	"github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api/resources/card"
	customer "github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api/resources/customer"
	sub "github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api/resources/subscription"
	token "github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api/resources/token"
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


	LoadCardClient(&card.Wrapper{})
	LoadCustomerClient(&customer.Wrapper{})
	LoadSubClient(&sub.Wrapper{DynamoSubRepo:&repo.Wrapper{DC:dynamoClient}})
	LoadTokenClient(&token.Wrapper{})
}
