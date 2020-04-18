package configs

const (
	PORT = ":5001"

	/************** DynamoDB configs *************/
	EU_WEST_2         = "eu-west-2"
	UNIQUE_IDENTIFIER = "email"
	PW                = "password"
	PREMIUM           = "premium"
	APPLICATIONS      = "applications"
	TABLE_NAME        = "users"
	/*********************************************/
	/************** RabbitMQ configs *************/
	FANOUT_EXCHANGE = "accounts.fanout"
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
)
