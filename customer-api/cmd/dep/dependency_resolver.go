package dep

import (
	"github.com/ProjectReferral/Get-me-in/customer-api/internal/api/email"
	email_builder "github.com/ProjectReferral/Get-me-in/customer-api/internal/event-driven"
	"log"
)

//methods that are implemented on util
//and will be used
type ConfigBuilder interface{
	LoadEnvConfigs()
}

//internal specific configs are loaded at runtime
//takes in a object(implemented interface) of type ServiceConfigs
func Inject(builder ConfigBuilder) {

	//load the env into the object
	builder.LoadEnvConfigs()

	LoadEmailRepo(&email_builder.EmailStruct{})
}
//variable injected with the interface methods
func LoadEmailRepo (r email_builder.EmailBuilder){
	log.Println("Injecting Email Repo")
	email.Service = r
}
