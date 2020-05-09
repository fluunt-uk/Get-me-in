package models

import "github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api/resources/models"

type Subscriber struct {
	Customer 		models.Customer 	`json:"customer"`
	PaymentDetails 	models.CardDetails 	`json:"card_details"`
}

//used to update user details
type ChangeRequest struct {
	Field		string 	`json:"field"`
	NewBool		bool 	`json:"new_bool"`
	//type 1: single string value
	//type 2: map value
	//type 3: boolean value
	Type		int		`json:"type"`
}