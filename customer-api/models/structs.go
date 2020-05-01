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
	Intro       string
	Outro		string
	Premium 	string
	Description string
	Price 		int
}

type IncomingNotificationDataStruct struct {
	Template 	string	`json:"template"`
	Email 		string 	`json:"email"`
	Firstname 	string 	`json:"firstname"`
	Surname 	string 	`json:"surname"`
}

type IncomingPaymentDataStruct struct {
	Template 	string	`json:"template"`
	Email 		string 	`json:"email"`
	Firstname 	string	`json:"firstname"`
	Surname 	string	`json:"surname"`
	Premium 	string	`json:"premium"`
	Description string	`json:"description"`
	Price 		int 	`json:"price"`
}

type IncomingActionDataStruct struct {
	Template 	string	`json:"template"`
	Email 		string 	`json:"email"`
	Firstname 	string 	`json:"firstname"`
	Surname 	string 	`json:"surname"`
	Accesscode 	string 	`json:"accesscode"`
}

//{
//"template":"reset-password",
//"email":"sharjeel50@hotmail.co.uk",
//"firstname":"Sharjeel",
//"surname":"Jan",
//"accesscode":"1234"
//}

const (
	CANCEL_SUBSCRIPTION		= "cancel-subscription-notification"
	NEW_USER_VERIFY 		= "new-user-verify-action"
	RESET_PASSWORD 			= "reset-password-action"
	CREATE_SUBSCRIPTION 	= "create-subscription-notification"
	PAYMENT_INVOICE 		= "payment-invoice-payment"
	PAYMENT_CONFIRMATION 	= "payment-confirmation-payment"
	REMINDER 				= "reminder-notification"
	REFEREE_APPLICATION 	= "referee-application"
)
