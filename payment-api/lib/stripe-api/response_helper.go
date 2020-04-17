package stripe_api

import (
	"encoding/json"
	"github.com/ProjectReferral/Get-me-in/payment-api/configs"
	"net/http"
)

func ReturnSuccessJSON(w http.ResponseWriter, objKey string) {
	stripeObject := configs.StripeObjects[objKey]

	toString, err := json.Marshal(stripeObject)

	if !HandleError(err, w) {
		w.Write(toString)
		w.WriteHeader(http.StatusOK)
	}
}
