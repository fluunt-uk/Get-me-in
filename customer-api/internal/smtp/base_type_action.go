package smtp

import (
	t "github.com/ProjectReferral/Get-me-in/customer-api/lib/hermes/templates"
	s "github.com/ProjectReferral/Get-me-in/customer-api/models"
)

func BaseTypeActionEmail(typeof string, p s.IncomingActionDataStruct) (string, string) {

	switch typeof {

	case s.NEW_USER_VERIFY:

		return t.GenerateActionHTMLTemplate(p, s.ActionEmailStruct{
			Intro:       "Welcome to GMI! We're very excited to have you on board.",
			Instruct:    "To get started, please click here:",
			ButtonText:  "Confirm your account",
			ButtonColor: "#22BC66",
			Outro:       "Need help, or have questions? Just reply to this email, we'd love to help.",
		}), "This is the subject"


	case s.RESET_PASSWORD:

		return t.GenerateActionHTMLTemplate(p, s.ActionEmailStruct{
			Intro:       "You recently made a request to reset your password.",
			Instruct:    "Please click the link below to continue.",
			ButtonText:  "Reset Password",
			ButtonColor: "#fc2403",
			Outro:       "If you did not make this change or you believe an unauthorised person has accessed your account, go to {reset-password endpoint} to reset your password without delay.",
		}), "This is the subject"

	}

	return "",""
}