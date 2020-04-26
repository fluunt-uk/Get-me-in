package payment

import (
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentmethod"
	"net/http"
)

func CreatePayment(w http.ResponseWriter, r *http.Request) {
	params := &stripe.PaymentMethodParams{
		Card: &stripe.PaymentMethodCardParams{
			Number:   stripe.String("4242424242424242"),
			ExpMonth: stripe.String("4"),
			ExpYear:  stripe.String("2021"),
			CVC:      stripe.String("314"),
		},
		Type: stripe.String("card"),
	}
	pm, _ := paymentmethod.New(params)

	ReturnSuccessJSON(w, pm)

}
