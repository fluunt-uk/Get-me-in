package util

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/ProjectReferral/Get-me-in/pkg/dynamodb"
	"github.com/ProjectReferral/Get-me-in/pkg/security"
	"log"
	"os"
	"net/http"
)

//internal specific configs are loaded at runtime
func LoadEnvConfigs(connectionName string, name string, port string,
	searchParam string, genericModel interface{}) {

	env := os.Getenv("ENV")
	log.Printf("Environment: %s\n",env)
	log.Printf("Running on %s\n", port)
	dynamodb.GenericModel = genericModel
	dynamodb.SearchParam = searchParam

	switch env {
	case "UAT":
		dynamodb.DynamoTable = "uat-" + name
	case "PROD":
		dynamodb.DynamoTable = "prod-" + name
	default:
		dynamodb.DynamoTable = "dev-" + name
	}
	
	connectToDynamoDB(connectionName)
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
func connectToDynamoDB(n string) {

	c := credentials.NewSharedCredentials("", "default")

	err := dynamodb.Connect(c, n)

	if err != nil {
		panic(err)
	}
}