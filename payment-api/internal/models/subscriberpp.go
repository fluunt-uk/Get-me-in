package models

import "github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api/resources/models"

type Subscriber struct {
	Customer models.Customer `json:"customer"`
	PaymentDetails models.CardDetails `json:"card_details"`
}