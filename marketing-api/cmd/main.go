package main

import (
	"github.com/ProjectReferral/Get-me-in/marketing-api/cmd/dep"
	"github.com/ProjectReferral/Get-me-in/marketing-api/configs"
	"github.com/ProjectReferral/Get-me-in/marketing-api/internal/api"
	"github.com/ProjectReferral/Get-me-in/marketing-api/internal/models"
	"github.com/ProjectReferral/Get-me-in/util"
	"os"
)

func main() {

	//gets all the necessary configs into our object
	//completes connections
	//assigns connections to repos
	dep.Inject(&util.ServiceConfigs{
		Environment: 	os.Getenv("ENV"),
		Region:       	configs.EU_WEST_2,
		Table:        	configs.TABLE_NAME,
		SearchParam:  	configs.UNIQUE_IDENTIFIER,
		GenericModel: 	models.Advert{},
		Port:		  	configs.PORT,
		BrokerUrl: 		configs.QAPI_URL,
	})

	api.SetupEndpoints()
}
