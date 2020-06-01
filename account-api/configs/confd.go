package configs

const (
	PORT = ":5001"

	/************** DynamoDB configs *************/
	EU_WEST_2         = "eu-west-2"
	UNIQUE_IDENTIFIER = "email"
	PW                = "password"
	PREMIUM           = "premium"
	APPLICATIONS      = "applications"
	ACTIVE_SUB      = "active_subscription"
	TABLE_NAME        = "users"
	/*********************************************/
	/************** RabbitMQ configs *************/
	FANOUT_EXCHANGE = "accounts.fanout"
	//for dev usage outside of local network
	//QAPI_URL = "http://35.179.11.178:5004"
	QAPI_URL = "http://localhost:5004"
	/*********************************************/
	/*********** Authentication configs **********/
	AUTH_REGISTER      = "register_user"
	AUTH_AUTHENTICATED = "crud"
	AUTH_LOGIN         = "signin_user"
	AUTH_VERIFY        = "verify_user"
	NO_ACCESS          = "admin_gui"
	/*********************************************/
)

var (
	//To dial RabbitMQ
	BrokerUrl = ""
	Env = ""
)
