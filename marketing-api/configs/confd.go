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
)
