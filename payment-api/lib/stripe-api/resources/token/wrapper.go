package card

import (
	stripe_api "github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/token"
	"net/http"
)

type Builder interface {
	CreateToken(http.ResponseWriter, *http.Request)
	GetToken(http.ResponseWriter, *http.Request)
}

type Wrapper struct{}

func (cw *Wrapper) CreateToken(w http.ResponseWriter, r *http.Request)  {
	params := &stripe.TokenParams{
		Card: &stripe.CardParams{
			Number: stripe.String("4242424242424242"),
			ExpMonth: stripe.String("12"),
			ExpYear: stripe.String("2021"),
			CVC: stripe.String("123"),
		},
	}
	t, _ := token.New(params)

	stripe_api.ReturnSuccessJSON(w, &t)
}

func (cw *Wrapper) GetToken(w http.ResponseWriter, r *http.Request)  {
	t, _ := token.Get(
		"tok_1GUZNNGhy1brUyYInPwRWKkA",
		nil,
	)

	stripe_api.ReturnSuccessJSON(w, &t)
}
