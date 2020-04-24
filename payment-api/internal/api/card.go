package api

import (
	"github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api/resources/card"
	"net/http"
)

var Service card.Builder

func CreateCard(w http.ResponseWriter, r *http.Request)  {
	Service.CreateCard(w, r)
}

func GetCard(w http.ResponseWriter, r *http.Request)  {
	Service.GetCard(w, r)
}

func UpdateCard(w http.ResponseWriter, r *http.Request)  {
	Service.UpdateCard(w, r)
}

func DeleteCard(w http.ResponseWriter, r *http.Request)  {
	Service.DeleteCard(w, r)
}

func GetAllCards(w http.ResponseWriter, r *http.Request)  {
	Service.GetAllCards(w, r)
}