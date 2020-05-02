package configs


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
	/*********************************************/

	)

var (
	//TODO: this key will need to be changed
	StripeKey = "sk_test_IsdgFVzr2pBjUip3oAzMbI5r007L1vhtUs"
	BrokerUrl = ""
)
