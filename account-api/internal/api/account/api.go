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
	//dynamodb.TestUpdate()
	//dynamodb.UpdateSingleField("surname", "lunos@gmail.com", "just trying")
	UpdateUser(w,r)
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
		w.Write([]byte("No body error"))
		return
	}
	body := r.Body

	u.AccessCode = event.NewUUID()

	dynamoAttr, errDecode := dynamodb.DecodeToDynamoAttribute(body, &u)
	dynamodb.ParseEmptyCollection(dynamoAttr, configs.APPLICATIONS)

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

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var cr models.ChangeRequest

	dynamodb.DecodeToMap(r.Body, &cr)

	UpdateValue(r.Header.Get("email"), &cr)

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

func VerifyEmail(w http.ResponseWriter, r *http.Request) {
	queryMap := r.URL.Query()

	accessCodeKeys, ok := queryMap["access_code"]
	tokenKeys, ok := queryMap["token"]

	accessCodeValue := accessCodeKeys[0]
	tokenValue := tokenKeys[0]

	if !ok || len(accessCodeValue) < 1 || len(tokenValue) < 1 {
		w.Write([]byte("Url Param are missing"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//validate the token expiry date
	if security.VerifyTokenWithClaim(tokenValue, configs.AUTH_VERIFY) {

		//email parsed from the jwt
		email := security.GetClaimsOfJWT().Subject
		result, err := dynamodb.GetItem(email)

		if !HandleError(err, w) {
			if accessCodeValue == *result.Item["access_code"].S {
				w.WriteHeader(http.StatusOK)
				return
			}
		}
	}

	w.WriteHeader(http.StatusBadRequest)
}
