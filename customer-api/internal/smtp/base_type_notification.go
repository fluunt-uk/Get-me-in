package smtp

import (
	t "github.com/ProjectReferral/Get-me-in/customer-api/lib/hermes/templates"
	s "github.com/ProjectReferral/Get-me-in/customer-api/models"
)

func BaseTypeNotificationEmail(typeof string, d []byte) (string, string) {

	switch typeof {

	case s.CREATE_SUBSCRIPTION:

		p := s.IncomingNotificationDataStruct{}
		t.ToStruct(d, &p)

		return t.GenerateNotificationHTMLTemplate(p, s.NotificationEmailStruct{
			Intro: "Welcome! Your GMI experience just got premium.",
			Outro: "",
		}), "This is the subject"

	case s.CANCEL_SUBSCRIPTION:

		p := s.IncomingNotificationDataStruct{}
		t.ToStruct(d, &p)

		// Will need to pass through some button to link to reactivate account (if possible) or pass in a button for sign up
		// Will also need to pass through when the service ends for the customer
		return t.GenerateNotificationHTMLTemplate(p, s.NotificationEmailStruct{
			Intro: "This is a confirmation that your GMI account has been canceled at your request.",
			Outro: "To start applying again, you can reactivate your account at any time. We hope you decide to come back soon.",
		}), "This is the subject"



	case s.REFEREE_APPLICATION:

		p := s.IncomingNotificationDataStruct{}
		t.ToStruct(d, &p)

		// Will need to pass through some button to link to reactivate account (if possible) or pass in a button for sign up
		// Will also need to pass through when the service ends for the customer
		return t.GenerateNotificationHTMLTemplate(p, s.NotificationEmailStruct{
			Intro: "",
			Outro: "",
		}), "This is the subject"


	case s.REMINDER:

		p := s.IncomingNotificationDataStruct{}
		t.ToStruct(d, &p)

		return t.GenerateNotificationHTMLTemplate(p, s.NotificationEmailStruct{
			Intro: "",
			Outro: "",
		}), "This is the subject"

	}

	return "",""
}
