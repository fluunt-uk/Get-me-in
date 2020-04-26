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

const (
	CANCEL_SUBSCRIPTION 	= "cancel-subscription"
	NEW_USER_VERIFY 		= "new-user-verify"
	RESET_PASSWORD 			= "reset-password"
	CREATE_SUBSCRIPTION 	= "create-subscription"
	PAYMENT_INVOICE 		= "payment-invoice"
	PAYMENT_CONFIRMATION 	= "payment-confirmation"
	REMINDER 				= "reminder"
	REFEREE_APPLICATION 	= "referee-application"
)
