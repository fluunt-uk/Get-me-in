package dep

import (
	"github.com/ProjectReferral/Get-me-in/payment-api/internal/api"
	"github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api/resources/card"
)

func LoadCardService(b card.Builder){
	api.StripeClient = b
}
