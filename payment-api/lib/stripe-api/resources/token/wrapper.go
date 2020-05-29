package card

import (
	stripe_api "github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api"
	"github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api/resources/models"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/token"
	"net/http"
)

type Builder interface {
	Put(details models.CardDetails) (*stripe.Token, error)
	Get(http.ResponseWriter, *http.Request)
}

type Wrapper struct{}

func (cw *Wrapper) Put(m models.CardDetails) (*stripe.Token, error) {
	params := &stripe.TokenParams{
		Card: &stripe.CardParams{
			Number: stripe.String(m.Number),
			ExpMonth: stripe.String(m.ExpMonth),
			ExpYear: stripe.String(m.ExpYear),
			CVC: stripe.String(m.CVC),
		},
	}
	t, err := token.New(params)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func (cw *Wrapper) Get(w http.ResponseWriter, r *http.Request)  {
	t, _ := token.Get(
		"tok_1GUZNNGhy1brUyYInPwRWKkA",
		nil,
	)

	stripe_api.ReturnSuccessJSON(w, &t)
}
