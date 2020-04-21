package repo_builder

import (
	"encoding/json"
	"github.com/ProjectReferral/Get-me-in/marketing-api/internal/models"
	"github.com/ProjectReferral/Get-me-in/pkg/dynamodb"
	"net/http"
)

type AdvertWrapper struct {
	//dynamo client
	DC		*dynamodb.Wrapper
}
//implement only the necessary methods for each repository
//available to be consumed by the API
type AdvertBuilder interface{
	GetAdvert(http.ResponseWriter, *http.Request)
	UpdateAdvert(http.ResponseWriter, *http.Request)
	CreateAdvert(http.ResponseWriter, *http.Request)
	DeleteAdvert(http.ResponseWriter, *http.Request)
}
//interface with the implemented methods will be injected in this variable
var Advert AdvertBuilder

//get all the adverts for a specific account
//token validated


//We check for the recaptcha response and proceed
//Covert the response body into appropriate models
//Create a new user using our dynamodb adapter
//A event message it sent to the queues which are consumed by the relevant services
func (a *AdvertWrapper) CreateAdvert(w http.ResponseWriter, r *http.Request) {

	dynamoAttr, errDecode := dynamodb.DecodeToDynamoAttribute(r.Body, models.Advert{})

	if !HandleError(errDecode, w, false) {

		err := dynamodb.CreateItem(dynamoAttr)

		if !HandleError(err, w, false) {
			//TODO: send event to queue
			w.WriteHeader(http.StatusOK)
		}
	}
}

func (a *AdvertWrapper) DeleteAdvert(w http.ResponseWriter, r *http.Request) {

	extractValue := ExtractValue(w, r)

	errDelete := dynamodb.DeleteItem(extractValue)

	if !HandleError(errDelete, w, false) {

		//Check item still exists
		result, err := dynamodb.GetItem(extractValue)

		//error thrown, record not found
		if !HandleError(err, w, true) {
			http.Error(w, result.GoString(), 302)
		}
	}
}

func (a *AdvertWrapper) GetAdvert(w http.ResponseWriter, r *http.Request) {

	result, err := dynamodb.GetItem(ExtractValue(w, r))

	if !HandleError(err, w, true) {
		b, err := json.Marshal(dynamodb.Unmarshal(result, models.Advert{}))

		if !HandleError(err, w, false) {

			w.Write([]byte(b))
			w.WriteHeader(http.StatusOK)
		}
	}
}

//Creating a new user with same ID replaces the record
//Temporary solution
func (a *AdvertWrapper) UpdateAdvert(w http.ResponseWriter, r *http.Request) {

	//TODO: Change to UpdateItem
	a.CreateAdvert(w, r)
}

//to avoid duplication, this method is re-used
//Gets the unique identifier from the response body
//This unique identifier is set under the API configs
//For this context, it would be id
//TODO: move to dynamodb library?
func ExtractValue(w http.ResponseWriter, r *http.Request) string {

	v, err := dynamodb.GetParameterValue(r.Body, models.Advert{})
	HandleError(err, w, false)

	return v
}
