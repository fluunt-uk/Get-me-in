package templates

import (
	"github.com/ProjectReferral/Get-me-in/customer-api/models"
	"github.com/matcornic/hermes"
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

	t := ActionEmail(models.ActionEmailStruct{
		Name:        k.Firstname,
		Intro:       l.Intro,
		Instruct:    l.Instruct,
		ButtonText:  l.ButtonText,
		ButtonColor: l.ButtonColor,
		ButtonLink:  "endpoint/" + k.Accesscode,
		Outro:       l.Outro,
	})

	return t
}