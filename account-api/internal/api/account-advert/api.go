package account_advert

//get the email from the jwt
//stored in the subject claim
//func GetAllAdverts(w http.ResponseWriter, r *http.Request) {
//	var u = models.User{}
//
//	//email parsed from the jwt
//	email := security.GetClaimsOfJWT().Subject
//	result, err := dynamodb.GetItem(email)
//
//	if !internal.HandleError(err, w) {
//
//		dynamodb.Unmarshal(result, &u)
//
//		b, err := json.Marshal(u.Applications)
//
//		if !internal.HandleError(err, w) {
//
//			w.Write(b)
//			w.WriteHeader(http.StatusOK)
//		}
//	}
//}
