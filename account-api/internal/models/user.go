package models

import event_driven "github.com/ProjectReferral/Get-me-in/account-api/internal/event-driven"

//'json:' is the value that will be picked up from the JSON body
//JSON must contain the value after 'json:...'  instead of the attribute name

type Entity interface {
	GenerateAccessCode()
}

type User struct {
	Uuid      	string `json:"id"`
	Firstname 	string `json:"firstname"`
	Surname   	string `json:"surname"`
	Email     	string `json:"email"`
	Password  	string `json:"password"`
	AccessCode	string `json:"accesscode"`
}

type Credentials struct {
	Email     string `json:"email"`
	Password  string `json:"password"`

}


func (u User) GenerateAccessCode() bool{

	u.AccessCode = event_driven.NewUUID()
	return true
}