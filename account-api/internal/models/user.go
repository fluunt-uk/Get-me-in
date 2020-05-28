package models

//'json:' is the value that will be picked up from the JSON body
//JSON must contain the value after 'json:...'  instead of the attribute name
type User struct {
	Uuid       		string 					`json:"id"`
	Firstname  		string 					`json:"first_name"`
	Surname    		string 					`json:"surname"`
	Email      		string 					`json:"email"`
	Password   		string 					`json:"password"`
	//for verify email purposes
	AccessCode 		string 					`json:"access_code"`
	//if the user has an active sub
	Premium    		bool 					`json:"premium"`
	//email has been verified
	Verified		bool					`dynamodbav:"verified"`
	//all the adverts the user has applied for
	Applications 	map[string]Advert 		`dynamodbav:"applications"`
	//active paid subscription
	ActiveSub	 	map[string]interface{}	`dynamodbav:"active_subscription"`
}

type Credentials struct {
	Email    		string 					`json:"email"`
	Password 		string 					`json:"password"`
}

type Advert struct {
	Uuid       		string 					`json:"id"`
	AccountId  		string 					`json:"account_id"`
	Title      		string 					`json:"title"`
	MaxUsers    	string 					`json:"max_users"`
	Premium     	bool   					`json:"premium"`
	ValidFrom   	string 					`json:"valid_from"`
	ValidTill   	string 					`json:"valid_till"`
	Company     	string 					`json:"company"`
	Description 	string 					`json:"description"`
}

//used to update user details
type ChangeRequest struct {
	Id				string					`json:"id"`
	NewString 		string 					`json:"new_value"`
	Field			string 					`json:"field"`
	NewMap			interface{} 			`json:"new_map"`
	NewBool			bool 					`json:"new_bool"`
	//type 1: single string value
	//type 2: map value
	//type 3: boolean value
	Type			int						`json:"type"`
}

