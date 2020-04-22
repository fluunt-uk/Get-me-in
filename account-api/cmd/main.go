package main

import (
	"github.com/ProjectReferral/Get-me-in/account-api/cmd/dep"
	"github.com/ProjectReferral/Get-me-in/account-api/configs"
	"github.com/ProjectReferral/Get-me-in/account-api/internal/api"
	"github.com/ProjectReferral/Get-me-in/account-api/internal/models"
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
		Table:        configs.TABLE_NAME,
		SearchParam:  configs.UNIQUE_IDENTIFIER,
		GenericModel: models.User{},
		BrokerUrl:    os.Getenv("BROKERURL"),
		Port:		  configs.PORT,
	})

	api.SetupEndpoints()
}

