package card

import (
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/card"
	"net/http"
)

func CreateCard(w http.ResponseWriter, r *http.Request)  {
	params := &stripe.CardParams{
		Customer: stripe.String("cus_H7EDAZGzu81IFr"),
		Token: stripe.String("tok_1GYzmlGhy1brUyYIRzzU6nvl"),
	}
	c, _ := card.New(params)

	ReturnSuccessJSON(w, c)
}
