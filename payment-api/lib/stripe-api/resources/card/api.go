package card

import (
	"github.com/ProjectReferral/Get-me-in/payment-api/configs"
	stripe_api "github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/card"
	"net/http"
)

func CreateCard(w http.ResponseWriter, r *http.Request)  {
	params := &stripe.CardParams{
		Customer: stripe.String("cus_H7HyJY5cWLA7Uf"),
		Token: stripe.String("tok_1GZ3MzGhy1brUyYIYJiEpaZB"),
	}
	c, _ := card.New(params)

	configs.StripeObjects["card"] = c
	stripe_api.ReturnSuccessJSON(w,"card")
}

func GetCard(w http.ResponseWriter, r *http.Request)  {
	params := &stripe.CardParams{
		Customer: stripe.String("cus_H7EDAZGzu81IFr"),
	}
	c, _ := card.Get(
		"card_1GYzmlGhy1brUyYItScV9Lwo",
		params,
	)

	configs.StripeObjects["card"] = c
	stripe_api.ReturnSuccessJSON(w,"card")
}

func UpdateCard(w http.ResponseWriter, r *http.Request)  {
	params := &stripe.CardParams{
		Customer: stripe.String("cus_H7Dt44weDWU4s5"),
		Name: stripe.String("Jenny Rosen"),
	}
	c, _ := card.Update(
		"card_1GYzmlGhy1brUyYItScV9Lwo",
		params,
	)

	configs.StripeObjects["card"] = c
	stripe_api.ReturnSuccessJSON(w,"card")
}

func DeleteCard(w http.ResponseWriter, r *http.Request)  {
	params := &stripe.CardParams{
		Customer: stripe.String("cus_H7Dt44weDWU4s5"),
	}
	c, _ := card.Del(
		"card_1GYzmlGhy1brUyYItScV9Lwo",
		params,
	)

	configs.StripeObjects["card"] = c
	stripe_api.ReturnSuccessJSON(w,"card")
}

func GetAllCards(w http.ResponseWriter, r *http.Request)  {
	params := &stripe.CardListParams{
		Customer: stripe.String("cus_H7Dt44weDWU4s5"),
	}
	params.Filters.AddFilter("limit", "", "3")
	i := card.List(params)
	for i.Next() {
		c := i.Card()

		configs.StripeObjects["card"] = c
		stripe_api.ReturnSuccessJSON(w,"card")
	}

}