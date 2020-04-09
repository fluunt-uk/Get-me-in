package configs

const (
	PORT = ":5000"

	/*************** API ENDPOINTS **************/
	LOGIN_ENDPOINT = "http://localhost:5001/account/signin"
	//LOGIN_ENDPOINT = "http://localhost:5000/mock"
	/*********************************************/

	/*********** Authentication configs **********/
	SERVICE_ID = "auth"
	SUBJECT = "authentication"
	BEARER = "Bearer"
	//2.5 days
	EXPIRY = 3600
	//5 minutes
	TEMP_EXPIRY = 5
	AUTH_REGISTER = "register_user"
	AUTH_AUTHENTICATED = "crud"
	/*********************************************/
)

