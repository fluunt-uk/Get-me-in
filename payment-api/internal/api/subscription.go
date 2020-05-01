package api

import (
	sub "github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api/resources/subscription"
	"net/http"
)

var SubClient sub.Builder


func CreateSub(w http.ResponseWriter, r *http.Request)  {
	SubClient.Put(w, r)

}

func GetSub(w http.ResponseWriter, r *http.Request)  {
	SubClient.Get(w, r)
}

func UpdateSub(w http.ResponseWriter, r *http.Request)  {
	SubClient.Patch(w, r)
}

func CancelSub(w http.ResponseWriter, r *http.Request)  {
	SubClient.Cancel(w, r)
}

func GetAllSubs(w http.ResponseWriter, r *http.Request)  {
	SubClient.GetBatch(w, r)
}