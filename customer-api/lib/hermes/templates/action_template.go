package templates

import (
	"github.com/ProjectReferral/Get-me-in/customer-api/models"
	"github.com/matcornic/hermes"
	"strconv"
)

type EmailBuilder struct {
	st	map[string]*models.BaseEmail
}

func (aeb *EmailBuilder) Innit(){
	aeb.st = make(map[string]*models.BaseEmail)
}

func (aeb *EmailBuilder) AddStaticTemplate(key string, s *models.BaseEmail) {
	aeb.st[key] = s
}

// This will be used for two types of emails currently, reset password and email confirmation.
func (aeb *EmailBuilder) templateMapping(params models.BaseEmail) string {
	hermesTable := &hermes.Table{}
	hermesAction := &[]hermes.Action{}

	if params.Payment.Price != 0 {
		hermesTable = &hermes.Table{
			Data: [][]hermes.Entry{
				{
					{Key: "Premium", Value: params.Payment.Premium},
					{Key: "Description", Value: params.Payment.Description},
					{Key: "Price", Value: strconv.Itoa(params.Payment.Price)},
				},
			},
			Columns: hermes.Columns{
				// Custom style for each rows
				CustomWidth: map[string]string{
					"Premium":  "20%",
					"Price": "15%",
				},
				CustomAlignment: map[string]string{
					"Price": "right",
				},
			},
		}
	}

	if params.Action.ButtonColor != "" {
		hermesAction = &[]hermes.Action{
			{
				Instructions: params.Action.Instruct,
				Button: hermes.Button{
					Color: params.Action.ButtonColor,
					Text:  params.Action.ButtonText,
					Link:  params.Action.ButtonLink,
				},
			},
		}
	}

	email := hermes.Email{
		Body: hermes.Body{
			Name: params.Name,
			Intros: []string {
				params.Intro,
			},
			Outros: []string{
				params.Outro,
			},
			Table: *hermesTable,
			Actions: *hermesAction,
		},
	}

	return StringParsedHTML(email)
}

func (aeb *EmailBuilder) GenerateHTMLTemplate(k models.IncomingData) string {

	t := aeb.templateMapping(models.BaseEmail{
		Name:  			k.FirstName,
		Intro: 			aeb.st[k.Template].Intro,
		Outro: 			aeb.st[k.Template].Outro,
		Action: 		models.ActionEmail{
			Instruct:    	aeb.st[k.Template].Action.Instruct,
			ButtonText:  	aeb.st[k.Template].Action.ButtonText,
			ButtonColor: 	aeb.st[k.Template].Action.ButtonColor,
			ButtonLink:  	"account/verify?access_code=" + k.AccessCode + "&token=" + k.Token,
		},
		Payment: models.PaymentEmail{
			Premium:     	k.Payment.Premium,
			Description: 	k.Payment.Description,
			Price:       	k.Payment.Price,
		},
	})

	return t
}
