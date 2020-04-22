package auth

import (
	"encoding/json"
	"github.com/ProjectReferral/Get-me-in/auth-api/configs"
	"github.com/ProjectReferral/Get-me-in/auth-api/models"
	"github.com/ProjectReferral/Get-me-in/pkg/security"
	"io"
	"log"
	"time"
)


func IssueToken(expiry time.Duration, audience string, subject string, body io.ReadCloser) security.TokenResponse {
	t := time.Now()
	e := t.Add(expiry * time.Minute)
	var u models.UserResponse

	//assign the claims to our customer model
	token := &security.TokenClaims{
		Issuer:		configs.SERVICE_ID,
		Subject:	subject,
		//treat audience as scope(permissions the token has access to)
		Audience:   audience,
		IssuedAt:   t.Unix(),
		Expiration: e.Unix(),
		NotBefore:  t.Unix(),
		Id:         "NOT_SET",
	}

	if subject != "register" {
		errJson := json.NewDecoder(body).Decode(&u)

		if errJson != nil {
			log.Println("Error parsing data to UserResponse object")
		}
	}

	tr := security.TokenResponse{
		//GenerateToken is our security library
		AccessToken:	security.GenerateToken(token),
		TokenType:		configs.BEARER,
		ExpiresIn:		configs.EXPIRY,
		//No support for refresh tokens as of yet
		RefreshToken: 	"N/A",
		UserData:		u,
	}

	return tr
}

