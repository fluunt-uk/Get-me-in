package api

import (
	"github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api/resources/card"
	"net/http"
)

var CardClient card.Builder

func CreateCard(w http.ResponseWriter, r *http.Request)  {
	CardClient.Put(w, r)
}

func GetCard(w http.ResponseWriter, r *http.Request)  {
	CardClient.Get(w, r)
}

func UpdateCard(w http.ResponseWriter, r *http.Request)  {
	CardClient.Patch(w, r)
}

func DeleteCard(w http.ResponseWriter, r *http.Request)  {
	CardClient.Del(w, r)
}

func GetAllCards(w http.ResponseWriter, r *http.Request)  {
	CardClient.GetBatch(w, r)
}