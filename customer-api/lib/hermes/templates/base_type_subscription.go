package templates
//
//import (
//	s "github.com/ProjectReferral/Get-me-in/customer-api/models"
//)
//
//func BaseTypeSubscriptionEmail(typeof string, p s.IncomingPaymentDataStruct) (string, string) {
//
//	switch typeof {
//
//	case s.PAYMENT_CONFIRMATION:
//
//		return t.GenerateSubscriptionHTMLTemplate(p, s.PaymentEmailStruct{
//			Intro: "Your order has been processed successfully.",
//			Outro: "Thank you, enjoy your experience.",
//		}), "This is the subject"
//
//	case s.PAYMENT_INVOICE:
//
//		return t.GenerateSubscriptionHTMLTemplate(p, s.PaymentEmailStruct{}), "This is the subject"
//	}
//
//	return "",""
//}
//
