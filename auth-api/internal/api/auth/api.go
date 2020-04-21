package auth

import (
	"encoding/json"
	"fmt"
	"github.com/ProjectReferral/Get-me-in/auth-api/configs"
	request "github.com/ProjectReferral/Get-me-in/pkg/http_lib"
	"io/ioutil"
	"net/http"
)

//Validates the request as human/robot with recaptcha
//Validates the credentials via a request to the Account-API
//Token is issued as a JSON with an expiry time of 2.5days
//This token will allow the user to access the [/GET,/PATCH,/DELETE] endpoints for the Account-API
func VerifyCredentials(w http.ResponseWriter, req *http.Request) {

	//TODO: reCaptchacheck

	//empty body
	if req.ContentLength < 1 {
		http.Error(w, "Error parsing body", http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		http.Error(w, "Error parsing body", http.StatusBadRequest)
		return
	}

	//request to account api to verify credentials
	resp, errPost := request.Post(configs.LOGIN_ENDPOINT, body,
		map[string]string{configs.AUTHORIZATION: req.Header.Get(configs.AUTHORIZATION)})

	if errPost != nil {
		http.Error(w, errPost.Error(), 400)
		return
	}

	if resp.StatusCode != 200 {
		errorBody, errParse := ioutil.ReadAll(resp.Body)

		if errParse != nil {
			http.Error(w, "Error parsing body", http.StatusBadRequest)
			return
		}

		http.Error(w, string(errorBody), resp.StatusCode)
		return
	}

	//subject here is the email
	token := IssueToken(configs.EXPIRY, configs.AUTH_AUTHENTICATED, resp.Header.Get(configs.SUBJECT), resp.Body)

	b, err := json.Marshal(token)
	if err != nil {
		fmt.Sprintf(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

//A temporary token can be requested for registration
//This token will only allow the user to access the /PUT endpoint for the Account-API
func IssueRegistrationTempToken(w http.ResponseWriter, req *http.Request) {
	token := IssueToken(configs.TEMP_EXPIRY, configs.AUTH_REGISTER, "register", nil)

	b, err := json.Marshal(token)

	if err != nil {
		fmt.Sprintf(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

//Response for testing purposes
func MockResponse(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("OK"))
	w.WriteHeader(http.StatusOK)
}
