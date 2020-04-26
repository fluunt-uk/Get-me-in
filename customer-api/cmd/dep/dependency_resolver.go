package dep

import (
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

	LoadEmailRepo(&email_builder.EmailStruct{})

	//load the env into the object
	builder.LoadEnvConfigs()

}
//variable injected with the interface methods
func LoadEmailRepo (r email_builder.EmailBuilder){
	log.Println("Injecting Email Repo")
	email_builder.Emails = r
}
