package main

import (
	"github.com/ProjectReferral/Get-me-in/payment-api/cmd/dep"
	"github.com/ProjectReferral/Get-me-in/payment-api/configs"
	"github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api/resources/models"
	"github.com/ProjectReferral/Get-me-in/util"
	"log"
	"os"
)

func main() {

	f, err := os.OpenFile("logs/paymentAPI_log.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	log.SetOutput(f)

	//gets all the necessary configs into our object
	//completes connections
	//assigns connections to repos
	dep.Inject(&util.ServiceConfigs{
		Environment: os.Getenv("ENV"),
		Region:       configs.EU_WEST_2,
		Table:        configs.TABLE,
		SearchParam:  configs.UNIQUE_IDENTIFIER,
		GenericModel: models.Subscription{},
		BrokerUrl:    configs.QAPI_URL,
		Port:		  configs.PORT,
	})
}
