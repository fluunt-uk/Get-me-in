package models

type Subscription struct {
	Email 			string  `json:"email"`
	AccountID 		string	`json:"account_id"`
	SubscriptionID 	string	`json:"sub_id"`
	PlanID 			string	`json:"plan_id"`
	PlanType 		string	`json:"plan_type"`
	Template 		string	`json:"template"`
	Price 			int64 	`json:"price"`
}

func (s *Subscription) SetTemplate(t string){
	s.Template = t
}