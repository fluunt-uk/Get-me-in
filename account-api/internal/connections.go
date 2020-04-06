package internal

import (
	"github.com/ProjectReferral/Get-me-in/account-api/configs"
	"github.com/ProjectReferral/Get-me-in/pkg/dynamodb"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

//Create a single instance of DynamoDB connection
func ConnectToDynamoDB() {

	c := credentials.NewSharedCredentials("", "default")

	err := dynamodb.Connect(c, configs.EU_WEST_2)

	if err != nil {
		panic(err)
	}
}
