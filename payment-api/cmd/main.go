package main

import (
	"github.com/ProjectReferral/Get-me-in/payment-api/cmd/dep"
	"github.com/ProjectReferral/Get-me-in/payment-api/internal"
)

func main() {
	dep.Inject()
	internal.ConnectToDynamoDB()
	internal.SetupEndpoints()
}
