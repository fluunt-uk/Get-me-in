package configs


var (
	PORT = ":5005"
	DevEmail = "project181219@gmail.com"
	DevEmailPw = "kowbu1-nuQjik-zyxput"
)

const (
	EXPIRY = 15
	ISSUER = "customer-api"
	AUDIENCE = "verify_user"
	SUB_ACTION_EMAIL = "http://localhost:5005/email/action"
)