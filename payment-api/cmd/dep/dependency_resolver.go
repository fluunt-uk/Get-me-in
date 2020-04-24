package dep

import (
	"github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api/resources/card"
)

func Inject(){

	LoadCardService(&card.Wrapper{})
}

