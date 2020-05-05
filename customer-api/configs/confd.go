package configs

//these can be manipulated during runtime
var (

	/*********** SMTP configs **********/
	DevEmail = "project181219@gmail.com"
	DevEmailPw = "kowbu1-nuQjik-zyxput"
	/*********************************************/
)

//these cannot be manipulated during runtime
const (
	PORT = ":5005"

	/*********** Authentication configs **********/
	EXPIRY = 15
	ISSUER = "customer-api"
	AUDIENCE = "verify_user"
	/*********************************************/

	/************** RabbitMQ configs *************/
	EMAIL_DISPATCHER_URL = "http://localhost:5005/smtp/send"
	QAPI_URL = "http://localhost:5004"
	//QAPI_URL = "http://35.179.11.178:5004"
	VERIFY_EMAIL_Q = "new-user-verify-email"
	/*********************************************/

	/***************** Email Template configs ****************/
	CANCEL_SUBSCRIPTION		= "cancel-subscription-notification"
	NEW_USER_VERIFY 		= "new-user-verify"
	RESET_PASSWORD 			= "reset-password"
	CREATE_SUBSCRIPTION 	= "create-subscription-notification"
	PAYMENT_INVOICE 		= "payment-invoice-payment"
	PAYMENT_CONFIRMATION 	= "payment-confirmation-payment"
	REMINDER 				= "reminder-notification"
	REFEREE_APPLICATION 	= "referee-application"
	/********************************************************/


)


