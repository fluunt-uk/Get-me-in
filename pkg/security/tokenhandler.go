package security

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

var JWTClaims = &jwt.StandardClaims{}

func GenerateToken(claim *TokenClaims) string {

	claims := &jwt.StandardClaims{
		Audience:  claim.Audience,
		ExpiresAt: claim.Expiration,
		Id:        claim.Id,
		IssuedAt:  claim.IssuedAt,
		Issuer:    claim.Issuer,
		NotBefore: claim.NotBefore,
		Subject:   claim.Subject,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	// TODO: use key from env
	tokenString, err := token.SignedString([]byte("this is the sample key"))

	if err != nil {
		return err.Error()
	}

	return tokenString
}

//Verify the token signature and expire dates without any explicit claims
func VerifyToken(tokenString string) bool {

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	token, err := jwt.ParseWithClaims(tokenString, JWTClaims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// TODO: use key from env
		return []byte("this is the sample key"), nil
	})

	// token.valid checks for expiry date too on top of signature
	if token.Valid &&
		err == nil &&
		JWTClaims.Valid() == nil {
		return true
	}
	return false
}

//Verify the token signature and expire dates with a claim
func VerifyTokenWithClaim(tokenString string, claim string) bool {

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	token, err := jwt.ParseWithClaims(tokenString, JWTClaims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// TODO: use key from env
		return []byte("this is the sample key"), nil
	})

	// token.valid checks for expiry date too on top of signature
	if token.Valid &&
		JWTClaims.Valid() == nil &&
		JWTClaims.VerifyAudience(claim, true) &&
		err == nil {
		return true
	}
	return false
}

func GetClaimsOfJWT() *jwt.StandardClaims {
	return JWTClaims
}


//Parses the authentication token and validates against the @claim
//Some tokens can only authenticate with specific endpoints
func WrapHandlerWithSpecialAuth(handler http.HandlerFunc, claim string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		a := req.Header.Get("Authorization")

		//empty header
		if a == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("No Authorization JTW!!"))
			return
		}

		//not empty header and token is valid
		if a != "" {
			if claim == "" && VerifyToken(a) {
				handler(w, req)
				return
			} else if claim != "" && VerifyTokenWithClaim(a, claim) {
				handler(w, req)
				return
			}
		}

		//not empty header and token is invalid
		w.WriteHeader(http.StatusUnauthorized)
	}
}