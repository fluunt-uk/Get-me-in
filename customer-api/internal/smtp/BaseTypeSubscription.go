package smtp

import (
	s "github.com/ProjectReferral/Get-me-in/customer-api/models"
	"github.com/ProjectReferral/Get-me-in/customer-api/lib/hermes/templates"
)

func BaseTypeSubscriptionEmail(typeof string, d []byte) (string, string) {

	switch typeof {

	case s.PAYMENT_CONFIRMATION:

		p := s.IncomingPaymentDataStruct{}
		templates.ToStruct(d, &p)

		return templates.GenerateSubscriptionHTMLTemplate(p, s.PaymentEmailStruct{
			Intro: "Your order has been processed successfully.",
			Outro: "Thank you, enjoy your experience.",
		}), "This is the subject"


	case s.PAYMENT_INVOICE:

		p := s.IncomingPaymentDataStruct{}
		templates.ToStruct(d, &p)

		return templates.GenerateSubscriptionHTMLTemplate(p, s.PaymentEmailStruct{}), "This is the subject"
	}

	return "",""
}

