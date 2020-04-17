package token

import (
	"github.com/ProjectReferral/Get-me-in/payment-api/configs"
	stripe_api "github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api"
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

	configs.StripeObjects["token"] = t
	stripe_api.ReturnSuccessJSON(w,"token")
}

func GetToken(w http.ResponseWriter, r *http.Request)  {
	t, _ := token.Get(
		"tok_1GUZNNGhy1brUyYInPwRWKkA",
		nil,
	)

	configs.StripeObjects["token"] = t
	stripe_api.ReturnSuccessJSON(w,"token")
}
