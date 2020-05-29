package configs

const (
	PORT = ":5003"

	/************** DynamoDB configs *************/
	EU_WEST_2         = "eu-west-2"
	TABLE_NAME        = "adverts"
	UNIQUE_IDENTIFIER = "id"
	/*********************************************/
	/*********** Authentication configs **********/
	AUTH_REGISTER      = "new_advert"
	AUTH_AUTHENTICATED = "crud"
	NO_ACCESS          = "admin_gui"
	/*********************************************/
	/************** RabbitMQ configs *************/
	FANOUT_EXCHANGE = "advert.fanout"
	QAPI_URL = "http://35.179.11.178:5004"
	/*********************************************/
)
