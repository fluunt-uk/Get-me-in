package account

import (
	"encoding/json"
	"fmt"
	"github.com/ProjectReferral/Get-me-in/account-api/configs"
	event "github.com/ProjectReferral/Get-me-in/account-api/internal/event-driven"
	"github.com/ProjectReferral/Get-me-in/account-api/internal/models"
	"github.com/ProjectReferral/Get-me-in/pkg/dynamodb"
	"github.com/ProjectReferral/Get-me-in/pkg/security"
	"net/http"
)

func TestFunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

//We check for the recaptcha response and proceed
//Covert the response body into appropriate models
//Create a new user using our dynamodb adapter
//A event message it sent to the queues which are consumed by the relevant services
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var u models.User

	//TODO: reCaptcha check, 30ms average
	if r.ContentLength < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No body error!"))
		return
	}
	body := r.Body

	u.GenerateAccessCode()
	dynamoAttr, errDecode := dynamodb.DecodeToDynamoAttribute(body, &u)

	if !HandleError(errDecode, w) {

		err := dynamodb.CreateItem(dynamoAttr)

		if !HandleError(err, w) {

			b, err := json.Marshal(u)
			if err != nil {
				fmt.Sprintf(err.Error())
			}
			//JSON format of the newly created user
			w.Write(b)
			w.WriteHeader(http.StatusOK)

			//triggers email confirmation e-mail
			go event.BroadcastUserCreatedEvent(string(b))
		}
	}

}

func GetUser(w http.ResponseWriter, r *http.Request) {

	//email parsed from the jwt
	email := security.GetClaimsOfJWT().Subject
	result, err := dynamodb.GetItem(email)

	if !HandleError(err, w) {
		b, err := json.Marshal(dynamodb.Unmarshal(result, models.User{}))

		if !HandleError(err, w) {

			w.Write(b)
			w.WriteHeader(http.StatusOK)
		}
	}
}

//Creating a new user with same ID replaces the record
//Temporary solution
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	//TODO: Change to UpdateItem
	CreateUser(w, r)
}

//check if the user has an active subscription
func IsUserPremium(w http.ResponseWriter, r *http.Request) {
	//email parsed from the jwt
	email := security.GetClaimsOfJWT().Subject
	result, err := dynamodb.GetItem(email)

	p := result.Item[configs.PREMIUM].BOOL

	if !HandleError(err, w) && *p {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(204)
	return
}

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