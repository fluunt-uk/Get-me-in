package sign_in

import (
	"github.com/ProjectReferral/Get-me-in/account-api/lib/dynamodb/repo"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	repo.SignIn.Login(w, r)
}
