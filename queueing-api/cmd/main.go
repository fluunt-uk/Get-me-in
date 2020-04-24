package main

import (
	"github.com/ProjectReferral/Get-me-in/queueing-api/internal/api"
)

func main() {
	api.Init()
	api.SetupEndpoints()
}