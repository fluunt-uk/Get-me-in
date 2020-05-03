package stripe_api

import (
	"encoding/json"
	"net/http"
)

func ReturnSuccessJSON(w http.ResponseWriter, i interface{}) {
	toString, err := json.Marshal(&i)

	if !HandleError(err, w) {
		w.Write(toString)
		w.WriteHeader(http.StatusOK)
	}
}
