package main

import (
	"github.com/ProjectReferral/Get-me-in/marketing-api/configs"
	"github.com/ProjectReferral/Get-me-in/marketing-api/internal/api"
	"github.com/ProjectReferral/Get-me-in/marketing-api/internal/models"
	"github.com/ProjectReferral/Get-me-in/util"
)

func main() {
	//internal specific configs are loaded at runtime
	util.LoadEnvConfigs(configs.EU_WEST_2, configs.TABLE_NAME, configs.PORT, configs.UNIQUE_IDENTIFIER, models.Advert{})
	api.SetupEndpoints()
}