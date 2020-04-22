package repo_builder

import (
	"encoding/json"
	"github.com/ProjectReferral/Get-me-in/marketing-api/internal/models"
	"github.com/ProjectReferral/Get-me-in/pkg/dynamodb"
	"net/http"
	"net/url"
)

type AdvertWrapper struct {
	//dynamo client
	DC *dynamodb.Wrapper
}

//implement only the necessary methods for each repository
//available to be consumed by the API
type AdvertBuilder interface {
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

		err := a.DC.CreateItem(dynamoAttr)

		if !HandleError(err, w, false) {
			//TODO: send event to queue
			w.WriteHeader(http.StatusOK)
		}
	}
}

func (a *AdvertWrapper) DeleteAdvert(w http.ResponseWriter, r *http.Request) {
	var am models.Advert

	//TODO:perhaps better to get from query string
	dynamodb.DecodeToMap(r.Body, &am)

	errDelete := a.DC.DeleteItem(am.Uuid)

	if !HandleError(errDelete, w, false) {

		//Check item still exists
		result, err := a.DC.GetItem(am.Uuid)

		//error thrown, record not found
		if !HandleError(err, w, true) {
			http.Error(w, result.GoString(), 302)
		}
	}
}

func (a *AdvertWrapper) GetAdvert(w http.ResponseWriter, r *http.Request) {
	var am models.Advert

	dynamodb.DecodeToMap(r.Body, &am)

	//TODO:perhaps better to get from query string
	result, err := a.DC.GetItem(am.Uuid)

	if !HandleError(err, w, true) {
		dynamodb.Unmarshal(result, &am)

		b, err := json.Marshal(&am)

		if !HandleError(err, w, false) {

			w.Write(b)
			w.WriteHeader(http.StatusOK)
		}
	}
}

//Creating a new user with same ID replaces the record
//Temporary solution
func (a *AdvertWrapper) UpdateAdvert(w http.ResponseWriter, r *http.Request) {
	var cr models.ChangeRequest

	advertID := GetQueryString(r.URL.Query(), "id", w)

	if advertID != "" {

		err := a.UpdateValue(advertID, &cr)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		w.WriteHeader(http.StatusOK)
	}
}

func GetQueryString(m url.Values, q string, w http.ResponseWriter) string {

	idKeys, ok := m[q]

	if !ok {
		w.Write([]byte("Url Param are missing"))
		w.WriteHeader(http.StatusBadRequest)
		return ""
	}

	advertID := idKeys[0]
	if len(advertID) < 1 {
		w.Write([]byte("Url Param are missing"))
		w.WriteHeader(http.StatusBadRequest)
		return ""
	}

	return advertID
}
