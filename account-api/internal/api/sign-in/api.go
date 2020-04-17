package sign_in

import (
	"encoding/json"
	"github.com/ProjectReferral/Get-me-in/account-api/internal/api/account"
	"github.com/ProjectReferral/Get-me-in/account-api/internal/models"
	"github.com/ProjectReferral/Get-me-in/pkg/dynamodb"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
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
	result, error := dynamodb.GetItem(u.Email)

	// if there is an error or record not found
	if error != nil {
		account.HandleError(error, w)
		return
	}
	var c models.Credentials

	dynamodb.Unmarshal(result, &c)
	passwordFromDB := c.Password

	//validation, hash matches
	if u.Password == passwordFromDB {
		w.Header().Add("subject", u.Email)
		b, err := json.Marshal(u)

		if !account.HandleError(err, w) {

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