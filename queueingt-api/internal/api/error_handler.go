package api

import (
	"github.com/ProjectReferral/Get-me-in/pkg/dynamodb"
	"net/http"
	"log"
)

//Custom made error
func HandleError(err error, w http.ResponseWriter) bool {
	if err != nil {
		//we can return the error with specific response code and reason
		e, isCustom := err.(*dynamodb.ErrorString)

		if isCustom {
			log.Println(e.Reason)
			w.WriteHeader(e.Code)
			return true
		}

		//default error
		log.Println(err.Error())
		w.WriteHeader(400)
		return true
	}
	return false
}
