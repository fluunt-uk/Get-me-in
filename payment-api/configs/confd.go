package configs


const (
	PORT = ":5003"

	/************** DynamoDB configs *************/
	EU_WEST_2         = "eu-west-2"
	UNIQUE_IDENTIFIER = "email"
	/*********************************************/)

var (
	//TODO: this key will need to be changed
	StripeKey = "sk_test_IsdgFVzr2pBjUip3oAzMbI5r007L1vhtUs"
	BrokerUrl = ""
)

var StripeObjects = make(map[string]interface{})
