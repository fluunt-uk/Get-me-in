package main

import (
	"github.com/ProjectReferral/Get-me-in/customer-api/cmd/dep"
	"github.com/ProjectReferral/Get-me-in/customer-api/configs"
	"github.com/ProjectReferral/Get-me-in/customer-api/internal/api"
	"github.com/ProjectReferral/Get-me-in/util"
	"log"
	"os"
)

func main() {

	dep.Inject(&util.ServiceConfigs{
		Environment: os.Getenv("ENV"),
		Port:		  configs.PORT,
	})

	loadEnvConfigs()
	api.SetupEndpoints()
}


// Will need to add email stuff to ServiceConfigs struct
func loadEnvConfigs() {
	log.Println("Service now running")

	configs.DevEmail = os.Getenv("DEVMAIL")
	configs.DevEmailPw = os.Getenv("DEVEMAILPW")

	env := os.Getenv("ENV")
	log.Println("Environment:" + env)
}