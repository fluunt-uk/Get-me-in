package main

import (
	"github.com/ProjectReferral/Get-me-in/customer-api/configs"
	"github.com/ProjectReferral/Get-me-in/customer-api/internal/api"
	"log"
	"os"
)

func main() {
	loadEnvConfigs()
	api.SetupEndpoints()
}


//internal specific configs are loaded at runtime
func loadEnvConfigs() {
	log.Println("Service now running")

	configs.DevEmail = os.Getenv("DEVMAIL")
	configs.DevEmailPw = os.Getenv("DEVEMAILPW")

	env := os.Getenv("ENV")
	log.Println("Environment:" + env)
}