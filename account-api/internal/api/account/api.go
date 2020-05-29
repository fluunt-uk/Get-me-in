package account

import (
	"github.com/ProjectReferral/Get-me-in/account-api/lib/dynamodb/repo-builder"
	"net/http"
)

func TestFunc(w http.ResponseWriter, r *http.Request) {
    	w.WriteHeader(http.StatusOK)
    	CreateUser(w,r)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	repo_builder.Account.CreateUser(w, r)
}

//get the email from the jwt
//stored in the subject claim
func GetUser(w http.ResponseWriter, r *http.Request) {
	repo_builder.Account.GetUser(w, r)
}

//two ways of updating a user's information
//type 1: updates a single string value for a defined field
//type 2: appends a map for a defined field(this field name must already exists)
//all parameters are set under ChangeRequest struct
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	repo_builder.Account.UpdateUser(w, r)
}

//check if the user has an active subscription
//parses email from the jwt
func IsUserPremium(w http.ResponseWriter, r *http.Request) {
	repo_builder.Account.IsUserPremium(w, r)
}

//we parse the access_code and token from the query string
//token is validated
//we compare the access_code in the db matches the one passed in from the query string
func VerifyEmail(w http.ResponseWriter, r *http.Request) {
	repo_builder.Account.VerifyEmail(w, r)
}

func ResendVerification(w http.ResponseWriter, r *http.Request) {
	repo_builder.Account.ResendVerification(w ,r)
}
