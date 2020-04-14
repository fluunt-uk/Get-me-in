package models

//'json:' is the value that will be picked up from the JSON body
//JSON must contain the value after 'json:...'  instead of the attribute name
type User struct {
	Uuid       string `json:"id"`
	Firstname  string `json:"firstname"`
	Surname    string `json:"surname"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	AccessCode string `json:"accesscode"`
	Premium    bool `json:"premium"`
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

