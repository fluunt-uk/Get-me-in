package main

import (
	"github.com/ProjectReferral/Get-me-in/account-api/configs"
	"github.com/ProjectReferral/Get-me-in/account-api/internal/api"
	"github.com/ProjectReferral/Get-me-in/account-api/internal/api/account"
	event_driven "github.com/ProjectReferral/Get-me-in/account-api/internal/event-driven"
	"github.com/ProjectReferral/Get-me-in/account-api/internal/models"
	"github.com/ProjectReferral/Get-me-in/account-api/lib/dynamodb/repo"
	"github.com/ProjectReferral/Get-me-in/pkg/dynamodb"
	"github.com/ProjectReferral/Get-me-in/util"
	"os"
)

func main() {
	DependencyInjection(&util.ServiceConfigs{
		Environment: os.Getenv("ENV"),
		Region:       configs.EU_WEST_2,
		Table:        configs.TABLE_NAME,
		SearchParam:  configs.UNIQUE_IDENTIFIER,
		GenericModel: models.User{},
		BrokerUrl:    os.Getenv("BROKERURL"),
		Port:		  configs.PORT,
	})

	api.SetupEndpoints()
}

//internal specific configs are loaded at runtime
func DependencyInjection(loader Loader) {

	loader.LoadEnvConfigs()
	//setup dynamo library
	dynamoClient := loader.LoadDynamoDBConfigs()
	//connect to the instance
	dynamoClient.Connect()

	//dependency injection to our resource
	//we inject the dynamo client
	account.Repo = &repo.DynamoLib{
		DC: dynamoClient,
	}
	//dependency injection to our resource
	//we inject the rabbitmq client
	event_driven.MQ = &event_driven.RabbitClient{
		URL:        os.Getenv("BROKERURL"),
	}

}

type Loader interface{
	LoadEnvConfigs()
	LoadDynamoDBConfigs() *dynamodb.DynamoDB
}