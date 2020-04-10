package main

import (
	"fmt"
	"github.com/ProjectReferral/Get-me-in/marketing-api/configs"
	"github.com/ProjectReferral/Get-me-in/marketing-api/internal"
	"github.com/ProjectReferral/Get-me-in/pkg/dynamodb"
	"log"
	"net/http"
	"os"
)

func main() {
	loadEnvConfigs()
	log.Fatal(http.ListenAndServe(configs.PORT, internal.SetupEndpoints()))
}

//internal specific configs are loaded at runtime
func loadEnvConfigs() {

	fmt.Print("Running on ")

	dynamodb.SearchParam = configs.UNIQUE_IDENTIFIER

	switch env := os.Getenv("ENV"); env {
	case "DEV":
		dynamodb.DynamoTable = "dev-adverts"
		fmt.Println(env)
	case "UAT":
		dynamodb.DynamoTable = "uat-adverts"
		fmt.Println(env)
	case "PROD":
		dynamodb.DynamoTable = "prod-adverts"
		fmt.Println(env)

	default:
		dynamodb.DynamoTable = "dev-adverts"
		fmt.Println(env)
	}
}
