package models

//'json:' is the value that will be picked up from the JSON body
//JSON must contain the value after 'json:...'  instead of the attribute name
type User struct {
	Uuid       	string `json:"id"`
	Firstname  	string `json:"first_name"`
	Surname    	string `json:"surname"`
	Email      	string `json:"email"`
	Password   	string `json:"password"`
	AccessCode 	string `json:"access_code"`
	Premium    	bool `json:"premium"`
	VerifyToken string `json:"verify_token"`
	Verified	bool `json:"verified`
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

