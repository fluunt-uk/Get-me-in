package models

type ActionEmailStruct struct {
	Name 		string
	Intro 		string
	Instruct 	string
	ButtonText 	string
	ButtonColor string
	ButtonLink 	string
	Outro 		string
}

type NotificationEmailStruct struct {
	Name  		string
	Intro 		string
	Outro 		string
}

type PaymentEmailStruct struct {
	Firstname 	string
	Premium 	string
	Description string
	Price 		int
}

type IncomingNotificationDataStruct struct {
	Email 		string 	`json:"email"`
	Firstname 	string 	`json:"firstname"`
	Surname 	string 	`json:"surname"`
}

type IncomingPaymentDataStruct struct {
	Email 		string 	`json:"email"`
	Firstname 	string	`json:"firstname"`
	Surname 	string	`json:"surname"`
	Premium 	string	`json:"premium"`
	Description string	`json:"description"`
	Price 		int 	`json:"price"`
}

type IncomingActionDataStruct struct {
	Email 		string 	`json:"email"`
	Firstname 	string 	`json:"firstname"`
	Surname 	string 	`json:"surname"`
	Accesscode 	string 	`json:"accesscode"`
}

func (c *IncomingPaymentDataStruct) Fullname() string {
	return c.Firstname + " " + c.Surname
}
