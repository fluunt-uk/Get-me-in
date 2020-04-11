package main

import (
	"fmt"
	"github.com/ProjectReferral/Get-me-in/marketing-api/configs"
	"github.com/ProjectReferral/Get-me-in/marketing-api/internal/api"
	"log"
	"net/http"
)

func main() {
	fmt.Println("test")
	//loadEnvConfigs()
	log.Fatal(http.ListenAndServe(configs.PORT, api.SetupEndpoints()))
}

//internal specific configs are loaded at runtime
func loadEnvConfigs() {
	//utils.loadEnvConfigs("adverts", configs.PORT)
}
