package main

import (
	event_driven "github.com/ProjectReferral/Get-me-in/customer-service/internal/event-driven"
)

func main() {
	event_driven.ReceiveFromAllQs()
}
