package repo_builder

import (
	"encoding/json"
	"github.com/ProjectReferral/Get-me-in/account-api/internal"
	"github.com/ProjectReferral/Get-me-in/account-api/internal/models"
	"github.com/ProjectReferral/Get-me-in/pkg/dynamodb"
	"github.com/ProjectReferral/Get-me-in/pkg/security"
	"net/http"
)
type AccountAdvertWrapper struct {
	//dynamo client
	DC		*dynamodb.Wrapper
}
//implement only the necessary methods for each repository
//available to be consumed by the API
type AccountAdvertBuilder interface{
	GetAllAdverts(w http.ResponseWriter, r *http.Request)
}
//interface with the implemented methods will be injected in this variable
var AccountAdvert AccountAdvertBuilder

//get all the adverts for a specific account
//token validated
func (c *AccountAdvertWrapper) GetAllAdverts(w http.ResponseWriter, r *http.Request) {
	var u = models.User{}

	//email parsed from the jwt
	email := security.GetClaimsOfJWT().Subject
	result, err := 	c.DC.GetItem(email)

	if !internal.HandleError(err, w) {

		dynamodb.Unmarshal(result, &u)

		b, err := json.Marshal(u.Applications)

		if !internal.HandleError(err, w) {

			w.Write(b)
			w.WriteHeader(http.StatusOK)
		}
	}
}
