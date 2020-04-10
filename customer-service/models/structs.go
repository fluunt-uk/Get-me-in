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
	Name  string
	Intro string
	Outro string
}

type PaymentEmailStruct struct {
	Firstname string
	Premium string
	Description string
	Price string
}

type IncomingDataStruct struct {
	Firstname string
	Surname string
	Email string
}