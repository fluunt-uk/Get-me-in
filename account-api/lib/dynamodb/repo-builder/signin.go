package repo_builder

import (
	"encoding/json"
	"github.com/ProjectReferral/Get-me-in/account-api/internal"
	"github.com/ProjectReferral/Get-me-in/account-api/internal/models"
	"github.com/ProjectReferral/Get-me-in/pkg/dynamodb"
	"net/http"
)

type SignInWrapper struct {
	//dynamo client
	DC		*dynamodb.Wrapper
}
//implement only the necessary methods for each repository
//available to be consumed by the API
type SignInBuilder interface{
	Login(w http.ResponseWriter, r *http.Request)
}
//interface with the implemented methods will be injected in this variable
var SignIn SignInBuilder

//credentials extract from the body
//query the db with the email
//if this exists, get the pw hash and compare
func (c *SignInWrapper) Login(w http.ResponseWriter, r *http.Request) {
	var u models.User

	errJson := json.NewDecoder(r.Body).Decode(&u)

	if errJson != nil {
		http.Error(w, errJson.Error(), 400)
		return
	}

	if isEmpty(u.Email, u.Password) {
		http.Error(w, "Invalid input", 400)
		return
	}
	//response from dynamodb
	result, error := c.DC.GetItem(u.Email)

	// if there is an error or record not found
	if error != nil {
		internal.HandleError(error, w)
		return
	}
	var cr models.Credentials

	dynamodb.Unmarshal(result, &cr)
	passwordFromDB := cr.Password

	//validation, hash matches
	if u.Password == passwordFromDB {
		w.Header().Add("subject", u.Email)
		b, err := json.Marshal(u)

		if !internal.HandleError(err, w) {

			w.Write(b)
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("Invalid credentials"))
}

func isEmpty(a string, b string) bool {
	return a == "" || b == ""
}