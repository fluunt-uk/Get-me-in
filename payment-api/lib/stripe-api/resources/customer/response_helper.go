package customer

import (
	"encoding/json"
	"github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api"
	"github.com/stripe/stripe-go"
	"net/http"
)

func ReturnSuccessJSON(w http.ResponseWriter, c *stripe.Customer){
	toString, err := json.Marshal(c)

	if stripe_api.HandleError(err, w) {
		w.Write(toString)
		w.WriteHeader(http.StatusOK)
	}
}
