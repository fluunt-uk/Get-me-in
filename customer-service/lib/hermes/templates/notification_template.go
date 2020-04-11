package templates

import (
	"github.com/ProjectReferral/Get-me-in/customer-service/models"
	"github.com/matcornic/hermes"
)

// This email will be used to only notifiy a user without any actionss
func NotificationEmail(params models.NotificationEmailStruct) string {
	email := hermes.Email{
		Body: hermes.Body{
			Name: params.Name,
			Intros: []string{
				params.Intro,
			},
			Outros: []string{
				params.Outro,
			},
		},
	}

	return StringParsedHTML(email)
}

func GenerateNotificationHTMLTemplate(p models.IncomingNotificationDataStruct, l models.NotificationEmailStruct) (string, string) {

	t := NotificationEmail(models.NotificationEmailStruct{
		Name:  p.Firstname,
		Intro: l.Intro,
		Outro: l.Outro,
	})

	return t, p.Email
}
