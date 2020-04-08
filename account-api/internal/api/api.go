package api

import (
	"encoding/json"
	"fmt"
	"github.com/ProjectReferral/Get-me-in/account-api/configs"
	event "github.com/ProjectReferral/Get-me-in/account-api/internal/event-driven"
	"github.com/ProjectReferral/Get-me-in/account-api/internal/models"
	"github.com/ProjectReferral/Get-me-in/pkg/dynamodb"
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

	//TODO: reCaptcha check, 30ms average
	if r.ContentLength < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No body error!"))
		return
	}
	body := r.Body

	dynamoAttr, errDecode, json := dynamodb.DecodeToDynamoAttributeAndJson(body, models.User{})

	if !HandleError(errDecode, w) {

		err := dynamodb.CreateItem(dynamoAttr)

		if !HandleError(err, w) {
			//JSON format of the newly created user
			w.Write([]byte(json))
			w.WriteHeader(http.StatusOK)

			//triggers email confirmation e-mail
			go event.BroadcastUserCreatedEvent(json)
		}
	}

}

func GetUser(w http.ResponseWriter, r *http.Request) {
	result, err := dynamodb.GetItem(getEmailValue())

	if !HandleError(err, w) {
		b, err := json.Marshal(dynamodb.Unmarshal(result, models.User{}))

		if !HandleError(err, w) {

			w.Write(b)
			w.WriteHeader(http.StatusOK)
		}
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	extractValue := ExtractValue(w, r)

	//Check item still exists
	result, err := dynamodb.GetItem(extractValue)

	//error thrown, record not found
	if !HandleError(err, w) {

		errDelete := dynamodb.DeleteItem(extractValue)

		if !HandleError(errDelete, w) {

			http.Error(w, result.GoString(), 204)
		}
	}
}

//Creating a new user with same ID replaces the record
//Temporary solution
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	//TODO: Change to UpdateItem
	CreateUser(w, r)
}

func Login(w http.ResponseWriter, r *http.Request) {

	// convert the response body into a map
	bodyMap, err := dynamodb.DecodeToMap(r.Body, models.Credentials{})

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if bodyMap != nil {
		//get the values
		emailFromBody := StringFromMap(bodyMap, configs.UNIQUE_IDENTIFIER)
		passwordFromBody := StringFromMap(bodyMap, configs.PW)

		if isEmpty(emailFromBody, passwordFromBody) {
			http.Error(w, "Invalid input", 400)
			return
		}
		//response from dynamodb
		result, error := dynamodb.GetItem(emailFromBody)

		// if there is an error or record not found
		if error != nil {
			HandleError(error, w)
			return
		}

		u := dynamodb.Unmarshal(result, models.Credentials{})
		passwordFromDB := StringFromMap(u, configs.PW)

		//validation, hash matches
		if passwordFromBody == passwordFromDB {
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	w.WriteHeader(http.StatusUnauthorized)
}

//Get the string value of a kay
func StringFromMap(m map[string]interface{}, p string) string {
	value := m[p]

	if value != nil {
		return fmt.Sprintf("%v", m[p])
	}

	return ""
}

func isEmpty(a string, b string) bool {
	return a == "" || b == ""
}

//to avoid duplication, this method is re-used
//Gets the unique identifier from the response body
//This unique identifier is set under the API configs
//For this context, it would be mail
//TODO: move to dynamodb library?
func ExtractValue(w http.ResponseWriter, r *http.Request) string {

	v, err := dynamodb.GetParameterValue(r.Body, models.User{})
	HandleError(err, w)

	return v
}
