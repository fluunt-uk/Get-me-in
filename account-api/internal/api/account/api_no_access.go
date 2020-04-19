package account

import (
	"encoding/json"
	"github.com/ProjectReferral/Get-me-in/account-api/internal/models"
	"github.com/ProjectReferral/Get-me-in/pkg/dynamodb"
	"net/http"
)

//SUPER USER operation
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var u models.User

	errJson := json.NewDecoder(r.Body).Decode(&u)

	if errJson != nil {
		http.Error(w, errJson.Error(), 400)
		return
	}

	//Check item still exists
	result, err := dynamodb.GetItem(u.Email)

	//error thrown, record not found
	if !HandleError(err, w) {

		errDelete := dynamodb.DeleteItem(u.Email)

		if !HandleError(errDelete, w) {

			http.Error(w, result.GoString(), 204)
		}
	}
}
