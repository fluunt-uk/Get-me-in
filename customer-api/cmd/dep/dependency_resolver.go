package dep

import (
	email_builder "github.com/ProjectReferral/Get-me-in/customer-api/internal/event-driven"
	"log"
)

//variable injected with the interface methods
func LoadSignInRepo (r email_builder.EmailBuilder){
	log.Println("Injecting SignIn Repo")
	email_builder.Emails = r
}
