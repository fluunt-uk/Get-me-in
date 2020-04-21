package api

import (
	repo_builder "github.com/ProjectReferral/Get-me-in/marketing-api/lib/dynamodb/repo-builder"
	"net/http"
)

func TestFunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}


//We check for the recaptcha response and proceed
//Covert the response body into appropriate models
//Create a new user using our dynamodb adapter
//A event message it sent to the queues which are consumed by the relevant services
func CreateAdvert(w http.ResponseWriter, r *http.Request) {

	repo_builder.Advert.CreateAdvert(w, r)
}

func DeleteAdvert(w http.ResponseWriter, r *http.Request) {

	repo_builder.Advert.DeleteAdvert(w, r)

}

func GetAdvert(w http.ResponseWriter, r *http.Request) {
	repo_builder.Advert.GetAdvert(w, r)
}

//Creating a new user with same ID replaces the record
//Temporary solution
func UpdateAdvert(w http.ResponseWriter, r *http.Request) {
	repo_builder.Advert.UpdateAdvert(w, r)
}

