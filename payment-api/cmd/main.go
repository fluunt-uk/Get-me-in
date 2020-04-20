package main

import (
	configs "github.com/ProjectReferral/Get-me-in/payment-api/configs"
	"github.com/ProjectReferral/Get-me-in/payment-api/internal"
	"github.com/ProjectReferral/Get-me-in/pkg/dynamodb"
	"github.com/stripe/stripe-go"
	"log"
	"os"
)

func main() {
	loadEnvConfigs()

	internal.SetupEndpoints()
}

//internal specific configs are loaded at runtime
func loadEnvConfigs() {
	stripe.Key = configs.StripeKey
	var env = ""

	log.Println("Running on %s \n", configs.PORT)

	configs.BrokerUrl = os.Getenv("BROKERURL")
	dynamodb.SearchParam = configs.UNIQUE_IDENTIFIER

	switch env := os.Getenv("ENV"); env {
	case "DEV":
		dynamodb.DynamoTable = "dev-users"
	case "UAT":
		dynamodb.DynamoTable = "uat-users"
	case "PROD":
		dynamodb.DynamoTable = "prod-users"
	default:
		dynamodb.DynamoTable = "dev-users"
	}

	log.Println("Environment:" + env)
}