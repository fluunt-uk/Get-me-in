package main

import (
	"github.com/ProjectReferral/Get-me-in/account-api/configs"
	"github.com/ProjectReferral/Get-me-in/account-api/internal/api"
	"github.com/ProjectReferral/Get-me-in/account-api/internal/models"
	"github.com/ProjectReferral/Get-me-in/util"
	"os"
)

func main() {
	loadEnvConfigs()
	api.SetupEndpoints()
}

//internal specific configs are loaded at runtime
func loadEnvConfigs() {
	configs.BrokerUrl = os.Getenv("BROKERURL")
	util.LoadEnvConfigs(configs.EU_WEST_2, configs.TABLE_NAME, configs.PORT, configs.UNIQUE_IDENTIFIER, models.User{})
}