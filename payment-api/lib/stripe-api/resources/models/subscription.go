package models

type Subscription struct {
	Email 			string  `json:"email"`
	AccountID 		string	`json:"id"`
	SubscriptionID 	string	`json:"id"`
	PlanID 			string	`json:"id"`
	PlanType 		string	`json:"nickname"`
}