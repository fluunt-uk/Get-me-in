package templates

import (
	"github.com/ProjectReferral/Get-me-in/customer-service/models"
	"github.com/matcornic/hermes"
	"strconv"
)

func PaymentEmail(params models.PaymentEmailStruct) string {
	email := hermes.Email{
		Body: hermes.Body{
				Name: params.Firstname,
				Intros: []string {
					params.Intro,
				},
			Table: hermes.Table{
				Data: [][]hermes.Entry{
					{
						{Key: "Premium", Value: params.Premium},
						{Key: "Description", Value: params.Description},
						{Key: "Price", Value: strconv.Itoa(params.Price)},
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
			},
			Outros: []string{
				params.Outro,
			},
		},
	}

	return StringParsedHTML(email)
}

func GenerateSubscriptionHTMLTemplate(p models.IncomingPaymentDataStruct, l models.PaymentEmailStruct) (string, string) {

	t := PaymentEmail(models.PaymentEmailStruct{
		Firstname: p.Firstname,
		Intro: l.Intro,
		Outro: l.Outro,
		Premium: p.Premium,
		Description: p.Description,
		Price: p.Price,
	})

	return t, p.Email
}


//{
//"email": "sharjeel50@hotmail.co.uk",
//"firstname": "Sharjeel",
//"surname": "Jan",
//"premium": "this is it",
//"Description": "mad ting",
//"price": 7
//}