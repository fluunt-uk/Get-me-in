package models

//'json:' is the value that will be picked up from the JSON body
//JSON must contain the value after 'json:...'  instead of the attribute name
type User struct {
	Uuid       		string `json:"id"`
	Firstname  		string `json:"first_name"`
	Surname    		string `json:"surname"`
	Email      		string `json:"email"`
	Password   		string `json:"password"`
	AccessCode 		string `json:"access_code"`
	Premium    		bool `json:"premium"`
	VerifyToken		string `json:"verify_token"`
	Verified		bool `json:"verified`
	Applications 	map[string]Advert `dynamodbav:"applications"`
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Advert struct {
	Uuid        string `json:"id"`
	AccountId   string `json:"account_id"`
	Title       string `json:"title"`
	MaxUsers    string `json:"max_users"`
	Premium     bool   `json:"premium"`
	ValidFrom   string `json:"valid_from"`
	ValidTill   string `json:"valid_till"`
	Company     string `json:"company"`
	Description string `json:"description"`
}

type ChangeRequest struct {
	NewString 	string 	`json:"new_value"`
	Field		string 	`json:"field"`
	NewMap		Advert 	`json:"new_map"`
	Type		int		`json:"type"`
}