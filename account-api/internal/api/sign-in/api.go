package sign_in

import (
	"encoding/json"
	"fmt"
	"github.com/ProjectReferral/Get-me-in/account-api/configs"
	"github.com/ProjectReferral/Get-me-in/account-api/internal/api"
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
		api.HandleError(error, w)
		return
	}

	c := dynamodb.Unmarshal(result, models.Credentials{})
	_, passwordFromDB := CredentialsFromMap(c, configs.UNIQUE_IDENTIFIER, configs.PW)

	//validation, hash matches
	if u.Password == passwordFromDB {
		w.Header().Add("subject", u.Email)
		b, err := json.Marshal(u)

		if !api.HandleError(err, w) {

			w.Write(b)
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("Invalid credentials"))
}


//Get the string value of a key
func CredentialsFromMap(m map[string]interface{}, u string, p string) (string, string) {
	username := m[u]
	password := m[p]

	if username != nil && password != nil {
		return fmt.Sprintf("%v", m[u]), fmt.Sprintf("%v", m[p])
	}

	return "", ""
}

func isEmpty(a string, b string) bool {
	return a == "" || b == ""
}