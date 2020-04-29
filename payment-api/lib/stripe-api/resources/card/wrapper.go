package card

import (
	stripe_api "github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/card"
	"net/http"
)

type Builder interface {
	Put(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	Del(http.ResponseWriter, *http.Request)
	Patch(http.ResponseWriter, *http.Request)
	GetBatch(http.ResponseWriter, *http.Request)
}

type Wrapper struct{}

func (cw *Wrapper) Put(w http.ResponseWriter, r *http.Request)  {
	params := &stripe.CardParams{
		Customer: stripe.String("cus_H7HyJY5cWLA7Uf"),
		Token: stripe.String("tok_1GZ3MzGhy1brUyYIYJiEpaZB"),
	}
	c, _ := card.New(params)

	stripe_api.ReturnSuccessJSON(w, &c)
}

func (cw *Wrapper) Get(w http.ResponseWriter, r *http.Request)  {
	params := &stripe.CardParams{
		Customer: stripe.String("cus_H7EDAZGzu81IFr"),
	}
	c, _ := card.Get(
		"card_1GYzmlGhy1brUyYItScV9Lwo",
		params,
	)

	stripe_api.ReturnSuccessJSON(w, &c)
}

func (cw *Wrapper) Patch(w http.ResponseWriter, r *http.Request)  {
	params := &stripe.CardParams{
		Customer: stripe.String("cus_H7Dt44weDWU4s5"),
		Name: stripe.String("Jenny Rosen"),
	}
	c, _ := card.Update(
		"card_1GYzmlGhy1brUyYItScV9Lwo",
		params,
	)

	stripe_api.ReturnSuccessJSON(w, &c)
}

func (cw *Wrapper) Del(w http.ResponseWriter, r *http.Request)  {
	params := &stripe.CardParams{
		Customer: stripe.String("cus_H7Dt44weDWU4s5"),
	}
	c, _ := card.Del(
		"card_1GYzmlGhy1brUyYItScV9Lwo",
		params,
	)

	stripe_api.ReturnSuccessJSON(w, &c)
}

func (cw *Wrapper) GetBatch(w http.ResponseWriter, r *http.Request)  {
	params := &stripe.CardListParams{
		Customer: stripe.String("cus_H7Dt44weDWU4s5"),
	}
	params.Filters.AddFilter("limit", "", "3")
	i := card.List(params)
	for i.Next() {
		c := i.Card()

		stripe_api.ReturnSuccessJSON(w, &c)
	}
}