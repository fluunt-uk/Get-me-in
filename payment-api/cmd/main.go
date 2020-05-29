package main

import (
	"github.com/ProjectReferral/Get-me-in/payment-api/cmd/dep"
	"github.com/ProjectReferral/Get-me-in/payment-api/configs"
	"github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api/resources/models"
	"github.com/ProjectReferral/Get-me-in/util"
	"os"
)

func main() {
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
