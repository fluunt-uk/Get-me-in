package api

import (
	"github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api/resources/card"
	"net/http"
)

var StripeClient card.Builder

func CreateCard(w http.ResponseWriter, r *http.Request)  {
	StripeClient.CreateCard(w, r)
}

func GetCard(w http.ResponseWriter, r *http.Request)  {
	StripeClient.GetCard(w, r)
}

func UpdateCard(w http.ResponseWriter, r *http.Request)  {
	StripeClient.UpdateCard(w, r)
}

func DeleteCard(w http.ResponseWriter, r *http.Request)  {
	StripeClient.DeleteCard(w, r)
}

func GetAllCards(w http.ResponseWriter, r *http.Request)  {
	StripeClient.GetAllCards(w, r)
}