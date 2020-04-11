package main

import (
	"fmt"
	_ "github.com/ProjectReferral/Get-me-in/marketing-api/configs"
	_ "github.com/ProjectReferral/Get-me-in/marketing-api/internal"
	_ "github.com/ProjectReferral/Get-me-in/utils"
	_ "log"
	_ "net/http"
	_ "os"
)

func main() {
	fmt.Println("test")
	//loadEnvConfigs()
	//log.Fatal(http.ListenAndServe(configs.PORT, internal.SetupEndpoints()))
}

//internal specific configs are loaded at runtime
func loadEnvConfigs() {
	//utils.loadEnvConfigs("adverts", configs.PORT)
}
