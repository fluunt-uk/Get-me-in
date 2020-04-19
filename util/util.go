package util

import (
	"github.com/ProjectReferral/Get-me-in/account-api/configs"
	dynamo_lib "github.com/ProjectReferral/Get-me-in/pkg/dynamodb"
	"github.com/ProjectReferral/Get-me-in/pkg/security"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"log"
	"net/http"
)



//internal specific configs are loaded at runtime
func (sc *ServiceConfigs) LoadEnvConfigs() {

	log.Printf("Environment: %s\n",sc.Environment)
	log.Printf("Running on %s\n", sc.Port)
}

//Parses the authentication token and validates against the @claim
//Some tokens can only authenticate with specific endpoints
func WrapHandlerWithSpecialAuth(handler http.HandlerFunc, claim string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		a := req.Header.Get("Authorization")

		//empty header
		if a == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("No Authorization JTW!!"))
			return
		}

		//not empty header and token is valid
		if a != "" {
			if claim == "" && security.VerifyToken(a) {
				handler(w, req)
				return
			} else if claim != "" && security.VerifyTokenWithClaim(a, claim) {
				handler(w, req)
				return
			}
		}

		//not empty header and token is invalid
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//Create a single instance of DynamoDB connection
func (sc *ServiceConfigs)generateCredentials() *credentials.Credentials{

	c := credentials.NewSharedCredentials("", "default")

	return c
}

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

type ServiceConfigs struct {
	Environment		string
	Region			string
	Table			string
	SearchParam		string
	GenericModel	interface{}
	BrokerUrl		string
	Port			string
}