package main

import (
	"github.com/ProjectReferral/Get-me-in/customer-service/configs"
	"github.com/ProjectReferral/Get-me-in/customer-service/internal/event-driven"
	"log"
	"os"
)

func main() {
	loadEnvConfigs()
	event_driven.ReceiveFromAllQs()
}


//internal specific configs are loaded at runtime
func loadEnvConfigs() {
	log.Println("Service now running")

	//get the env variables
	configs.BrokerUrl = os.Getenv("BROKERURL")
	configs.DevEmail = os.Getenv("DEVMAIL")
	configs.DevEmailPw = os.Getenv("DEVEMAILPW")

	env := os.Getenv("ENV")
	log.Println("Environment:" + env)
}