package util

import (
	"github.com/ProjectReferral/Get-me-in/account-api/configs"
	dynamo_lib "github.com/ProjectReferral/Get-me-in/pkg/dynamodb"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"log"
)

type ServiceConfigs struct {
	Environment		string
	Region			string
	Table			string
	SearchParam		string
	GenericModel	interface{}
	BrokerUrl		string
	Port			string
}

//internal specific configs are loaded at runtime
func (sc *ServiceConfigs) LoadEnvConfigs() {

	log.Printf("Environment: %s\n",sc.Environment)
	log.Printf("Running on %s\n", sc.Port)
}

//DynamoDB configs
func (sc *ServiceConfigs) LoadDynamoDBConfigs() *dynamo_lib.DynamoDB{

	switch configs.Env {
	case "UAT":
		sc.Table = "uat-" + sc.Table
	case "PROD":
		sc.Table = "prod-" + sc.Table
	default:
		sc.Table = "dev-" + sc.Table
	}

	dynamoDBInstance := &dynamo_lib.DynamoDB{
		GenericModel:sc.GenericModel,
		SearchParam:sc.SearchParam,
		Table:sc.Table,
		Credentials:sc.generateCredentials(),
		Region: sc.Region,
	}
	return dynamoDBInstance
}

//Create a single instance of DynamoDB connection
func (sc *ServiceConfigs)generateCredentials() *credentials.Credentials{

	c := credentials.NewSharedCredentials("", "default")

	return c
}
