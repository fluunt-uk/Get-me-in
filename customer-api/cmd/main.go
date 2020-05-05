package main

import (
	"github.com/ProjectReferral/Get-me-in/customer-api/cmd/dep"
	"github.com/ProjectReferral/Get-me-in/customer-api/configs"
	"github.com/ProjectReferral/Get-me-in/util"
	"os"
)

func main() {

	dep.Inject(&util.ServiceConfigs{
		Environment: os.Getenv("ENV"),
		Port:        configs.PORT,
		BrokerUrl:   configs.QAPI_URL,
	})
}
