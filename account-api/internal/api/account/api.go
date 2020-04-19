package account

import (
	"github.com/ProjectReferral/Get-me-in/account-api/lib/dynamodb/repo"
	"net/http"
)

var Repo *repo.DynamoLib

func TestFunc(w http.ResponseWriter, r *http.Request) {
	CreateUser(w, r)
	w.WriteHeader(http.StatusOK)
}

func CreateUser(w http.ResponseWriter, r *http.Request){
	Repo.CreateUser(w, r)
}


//
////get the email from the jwt
////stored in the subject claim
//func GetUser(w http.ResponseWriter, r *http.Request) {
//
//	//email parsed from the jwt
//	email := security.GetClaimsOfJWT().Subject
//	result, err := dynamodb.GetItem(email)
//
//	if !internal.HandleError(err, w) {
//		b, err := json.Marshal(dynamodb.Unmarshal(result, models.User{}))
//
//		if !internal.HandleError(err, w) {
//
//			w.Write(b)
//			w.WriteHeader(http.StatusOK)
//		}
//	}
//}
//
////two ways of updating a user's information
////type 1: updates a single string value for a defined field
////type 2: appends a map for a defined field(this field name must already exists)
////all parameters are set under ChangeRequest struct
//func UpdateUser(w http.ResponseWriter, r *http.Request) {
//	var cr models.ChangeRequest
//
//	dynamodb.DecodeToMap(r.Body, &cr)
//
//	email := security.GetClaimsOfJWT().Subject
//	err := UpdateValue(email, &cr)
//
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//	}
//
//	w.WriteHeader(http.StatusOK)
//}
//
////check if the user has an active subscription
////parses email from the jwt
//func IsUserPremium(w http.ResponseWriter, r *http.Request) {
//	//email parsed from the jwt
//	email := security.GetClaimsOfJWT().Subject
//	result, err := dynamodb.GetItem(email)
//
//	p := result.Item[configs.PREMIUM].BOOL
//
//	if !internal.HandleError(err, w) && *p {
//		w.WriteHeader(http.StatusOK)
//		return
//	}
//
//	w.WriteHeader(204)
//	return
//}
//
////we parse the access_code and token from the query string
////token is validated
////we compare the access_code in the db matches the one passed in from the query string
//func VerifyEmail(w http.ResponseWriter, r *http.Request) {
//	queryMap := r.URL.Query()
//
//	accessCodeKeys, ok := queryMap["access_code"]
//	tokenKeys, ok := queryMap["token"]
//	if !ok {
//		w.Write([]byte("Url Param are missing"))
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//
//	accessCodeValue := accessCodeKeys[0]
//	tokenValue := tokenKeys[0]
//	if len(accessCodeValue) < 1 || len(tokenValue) < 1 {
//		w.Write([]byte("Url Param are missing"))
//		w.WriteHeader(http.StatusBadRequest)
//	}
//
//	//validate the token expiry date
//	if security.VerifyTokenWithClaim(tokenValue, configs.AUTH_VERIFY) {
//
//		//email parsed from the jwt
//		email := security.GetClaimsOfJWT().Subject
//		result, err := dynamodb.GetItem("lunos4@gmail.com")
//
//		if !internal.HandleError(err, w) {
//			if accessCodeValue == *result.Item["access_code"].S {
//
//				UpdateValue(email, &models.ChangeRequest{Field: "verified", NewBool: true, Type: 3})
//				w.WriteHeader(http.StatusOK)
//				return
//			}
//			w.WriteHeader(http.StatusUnauthorized)
//			w.Write([]byte("Access code does not match"))
//			return
//		}
//	}
//
//	w.WriteHeader(http.StatusBadRequest)
//}
//
////TODO:resend verification email
//func ResendVerification(w http.ResponseWriter, r *http.Request) {
//	var u models.User
//	email := security.GetClaimsOfJWT().Subject
//
//	//new access code generated
//	UpdateValue(email, &models.ChangeRequest{Field: "access_code", NewString: event.NewUUID(), Type: 1})
//
//	user, err := dynamodb.GetItem("lunos4@gmail.com")
//
//	if !internal.HandleError(err, w) {
//
//		dynamodb.Unmarshal(user, &u)
//		b, errM := json.Marshal(&u)
//
//		if !internal.HandleError(errM, w) {
//
//			w.Write(b)
//			w.WriteHeader(http.StatusOK)
//
//			go event.BroadcastUserCreatedEvent(string(b))
//		}
//	}
//}
