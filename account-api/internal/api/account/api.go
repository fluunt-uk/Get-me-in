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

//get the email from the jwt
//stored in the subject claim
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

//two ways of updating a user's information
//type 1: updates a single string value for a defined field
//type 2: appends a map for a defined field(this field name must already exists)
//all parameters are set under ChangeRequest struct
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var cr models.ChangeRequest

	dynamodb.DecodeToMap(r.Body, &cr)

	email := security.GetClaimsOfJWT().Subject
	err := UpdateValue(email, &cr)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
}

//check if the user has an active subscription
//parses email from the jwt
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

//we parse the access_code and token from the query string
//token is validated
//we compare the access_code in the db matches the one passed in from the query string
func VerifyEmail(w http.ResponseWriter, r *http.Request) {
	queryMap := r.URL.Query()

	accessCodeKeys, ok := queryMap["access_code"]
	tokenKeys, ok := queryMap["token"]
	if !ok {
		w.Write([]byte("Url Param are missing"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	accessCodeValue := accessCodeKeys[0]
	tokenValue := tokenKeys[0]
	if len(accessCodeValue) < 1 || len(tokenValue) < 1 {
		w.Write([]byte("Url Param are missing"))
		w.WriteHeader(http.StatusBadRequest)
	}

	//validate the token expiry date
	if security.VerifyTokenWithClaim(tokenValue, configs.AUTH_VERIFY) {

		//email parsed from the jwt
		email := security.GetClaimsOfJWT().Subject
		result, err := dynamodb.GetItem("lunos4@gmail.com")

		if !HandleError(err, w) {
			if accessCodeValue == *result.Item["access_code"].S {

				UpdateValue(email, &models.ChangeRequest{Field: "verified", NewBool: true, Type: 3})
				w.WriteHeader(http.StatusOK)
				return
			}
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Access code does not match"))
			return
		}
	}

	w.WriteHeader(http.StatusBadRequest)
}

//TODO:resend verification email
func ResendVerification(w http.ResponseWriter, r *http.Request) {
	var u models.User
	email := security.GetClaimsOfJWT().Subject

	//new access code generated
	UpdateValue(email, &models.ChangeRequest{Field: "access_code", NewString: event.NewUUID(), Type: 1})

	user, err := dynamodb.GetItem("lunos4@gmail.com")

	if !HandleError(err, w) {

		dynamodb.Unmarshal(user, &u)
		b, errM := json.Marshal(&u)

		if !HandleError(errM, w) {

			w.Write(b)
			w.WriteHeader(http.StatusOK)

			go event.BroadcastUserCreatedEvent(string(b))
		}
	}
}
