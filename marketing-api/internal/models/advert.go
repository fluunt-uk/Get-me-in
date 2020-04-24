package models

//'json:' is the value that will be picked up from the JSON body
//JSON must contain the value after 'json:...'  instead of the attribute name
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

//used to update user details
type ChangeRequest struct {
	NewString 	string 	`json:"new_value"`
	Field		string 	`json:"field"`
	NewMap		Advert 	`json:"new_map"`
	NewBool		bool 	`json:"new_bool"`
	//type 1: single string value
	//type 2: map value
	//type 3: boolean value
	Type		int		`json:"type"`
}

