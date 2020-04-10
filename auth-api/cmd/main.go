package main

import (
	"github.com/ProjectReferral/Get-me-in/auth-api/configs"
	"github.com/ProjectReferral/Get-me-in/auth-api/internal"
	"log"
)

func main() {
	log.Println("Running on %s", configs.PORT)
	internal.SetupEndpoints()
}
