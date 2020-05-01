package api

import (
	token "github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api/resources/token"
	"net/http"
)

var TokenClient token.Builder

func CreateToken(w http.ResponseWriter, r *http.Request)  {
	//TokenClient.Put(w, r)
}

func GetToken(w http.ResponseWriter, r *http.Request)  {
	TokenClient.Get(w, r)
}

