package configs

import "time"

const (
	PORT = ":5003"

	/************** DynamoDB configs *************/
	EU_WEST_2         	= "eu-west-2"
	UNIQUE_IDENTIFIER 	= "email"
	TABLE				= "subs"
	/*********************************************/
	/************** Payments configs *************/
	PREMIUM_PLAN         	= "plan_H4eVnOxhxYYZ7a"
	ACCOUNT_API_PREMIUM		= "http://localhost:5001/account"
	CANCEL_SUBSCRIPTION		= "cancel-subscription"
	CREATE_SUBSCRIPTION 	= "create-subscription"
	PAYMENT_INVOICE 		= "payment-invoice"
	PAYMENT_CONFIRMATION 	= "payment-confirmation"
	/*********************************************/
	/*********** Authorization configs ***********/
	AUTH_HEADER         	= "Authorization"
	AUTH_AUTHENTICATED = "crud"
	/*********************************************/
	/************** RabbitMQ configs *************/
	FANOUT_EXCHANGE = "adverts.fanout"
	//for dev usage outside of local network
	//QAPI_URL = "http://35.179.11.178:5004"
	QAPI_URL = "http://localhost:5004"
	/*********************************************/
	THROTTLE = 600 * time.Millisecond
	)

var (
	//TODO: this key will need to be changed
	StripeKey = "sk_test_IsdgFVzr2pBjUip3oAzMbI5r007L1vhtUs"
	BrokerUrl = ""
)
