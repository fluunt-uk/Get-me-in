package account_advert

import (
	"github.com/ProjectReferral/Get-me-in/account-api/lib/dynamodb/repo"
	"net/http"
)

//get the email from the jwt
//stored in the subject claim
func GetAllAdverts(w http.ResponseWriter, r *http.Request) {
	repo.AccountAdvert.GetAllAdverts(w ,r)
}
