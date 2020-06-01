package models

type ActionEmail struct {
	Instruct 		string
	ButtonText 		string
	ButtonColor 	string
	ButtonLink 		string
}

type PaymentEmail struct {
	Premium 		string
	Description 	string
	Price 			int
}

type BaseEmail struct {
	Subject 		string
	Name  			string
	Intro 			string
	Outro 			string
	Payment			PaymentEmail
	Action			ActionEmail
}

type IncomingData struct{
	Template 		string				`json:"template"`
	Email 			string 				`json:"email"`
	FirstName 		string 				`json:"first_name"`
	Surname 		string 				`json:"surname"`
	AccessCode 		string 				`json:"access_code"`
	Token 			string 				`json:"token"`
	Payment			IncomingPaymentData	`json:"payment"`

}

type IncomingPaymentData struct {
	Premium 		string				`json:"premium"`
	Description 	string				`json:"description"`
	Price 			int 				`json:"price"`
}


//{
//"template":"reset-password",
//"email":"sharjeel50@hotmail.co.uk",
//"firstname":"Sharjeel",
//"surname":"Jan",
//"accesscode":"1234"
//}


