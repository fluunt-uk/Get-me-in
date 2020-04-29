package dep

import (
	"github.com/ProjectReferral/Get-me-in/payment-api/internal/api"
	"github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api/resources/card"
	customer "github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api/resources/customer"
	sub "github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api/resources/subscription"
	token "github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api/resources/token"
	"log"
)

func LoadCardClient(b card.Builder){
	log.Println("Setting up Card Client")
	api.CardClient = b
}

func LoadCustomerClient(b customer.Builder){
	log.Println("Setting up Customer Client")
	api.CustomerClient = b
}

func LoadTokenClient(b token.Builder){
	log.Println("Setting up Token Client")
	api.TokenClient = b
}

func LoadSubClient(b sub.Builder){
	log.Println("Setting up Subscription Client")
	api.SubClient = b
}