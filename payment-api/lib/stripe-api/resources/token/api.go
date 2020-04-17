package token

import (
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/token"
	"net/http"
)

func CreateToken(w http.ResponseWriter, r *http.Request)  {
	params := &stripe.TokenParams{
		Card: &stripe.CardParams{
			Number: stripe.String("4242424242424242"),
			ExpMonth: stripe.String("12"),
			ExpYear: stripe.String("2021"),
			CVC: stripe.String("123"),
		},
	}
	t, _ := token.New(params)

	ReturnSuccessJSON(w, t)
}
