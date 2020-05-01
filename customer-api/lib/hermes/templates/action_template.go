package templates

import (
	"github.com/ProjectReferral/Get-me-in/customer-api/configs"
	"github.com/ProjectReferral/Get-me-in/customer-api/models"
	"github.com/ProjectReferral/Get-me-in/pkg/security"
	"github.com/matcornic/hermes"
	"time"
)

// This will be used for two types of emails currently, reset password and email confirmation.
func ActionEmail(params models.ActionEmailStruct) string {

	email := hermes.Email{
		Body: hermes.Body{
			Name: params.Name,
			Intros: []string {
				params.Intro,
			},
			Actions: []hermes.Action{
				{
					Instructions: params.Instruct,
					Button: hermes.Button{
						Color: params.ButtonColor,
						Text:  params.ButtonText,
						Link:  params.ButtonLink,
					},
				},
			},
			Outros: []string{
				params.Outro,
			},
		},
	}

	return StringParsedHTML(email)
}

func GenerateActionHTMLTemplate(k models.IncomingActionDataStruct, l models.ActionEmailStruct) string {

	y := time.Now()
	e := y.Add(configs.EXPIRY * time.Minute)

	tk := security.GenerateToken(&security.TokenClaims{
		Issuer:     configs.ISSUER,
		Subject:    k.Email,
		Audience:   configs.AUDIENCE,
		IssuedAt:   y.Unix(),
		Expiration: e.Unix(),
		NotBefore:  y.Unix(),
		Id:         "NOT_SET",
	})

	t := ActionEmail(models.ActionEmailStruct{
		Name:        k.Firstname,
		Intro:       l.Intro,
		Instruct:    l.Instruct,
		ButtonText:  l.ButtonText,
		ButtonColor: l.ButtonColor,
		ButtonLink:  "account/verify?access_code=" + k.Accesscode + "&token=" + tk,
		Outro:       l.Outro,
	})

	return t
}