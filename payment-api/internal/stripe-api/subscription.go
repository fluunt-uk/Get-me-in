package stripe_api

import (
	"encoding/json"
	"github.com/ProjectReferral/Get-me-in/payment-api/internal"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/sub"
	"net/http"
)

func CreateSub(w http.ResponseWriter, r *http.Request) {
	params := &stripe.SubscriptionParams{
		Customer: stripe.String("cus_H4dYk0sB1TmcGB"),
		Items: []*stripe.SubscriptionItemsParams{
			{
				Plan: stripe.String("plan_EeE4ns3bvb34ZP"),
			},
		},
	}
	s, _ := sub.New(params)
	toString, err := json.Marshal(s)

	if !internal.HandleError(err, w) {
		w.Write(toString)
		w.WriteHeader(http.StatusOK)
	}
}

func RetrieveSub(w http.ResponseWriter, r *http.Request) {
	s, _ := sub.Get("sub_36VrPHS2vVxJMq", nil)
	toString, err := json.Marshal(s)

	if !internal.HandleError(err, w) {
		w.Write(toString)
		w.WriteHeader(http.StatusOK)
	}
}

func UpdateSub(w http.ResponseWriter, r *http.Request) {
	params := &stripe.SubscriptionParams{}
	params.AddMetadata("order_id", "6735")
	s, _ := sub.Update("sub_36VrPHS2vVxJMq", params)
	toString, err := json.Marshal(s)

	if !internal.HandleError(err, w) {
		w.Write(toString)
		w.WriteHeader(http.StatusOK)
	}
}

func CancelSub(w http.ResponseWriter, r *http.Request) {
	s, _ := sub.Cancel("sub_36VrPHS2vVxJMq", nil)
	toString, err := json.Marshal(s)

	if !internal.HandleError(err, w) {
		w.Write(toString)
		w.WriteHeader(http.StatusOK)
	}
}

func ListSubs(w http.ResponseWriter, r *http.Request) {
	params := &stripe.SubscriptionListParams{}
	params.Filters.AddFilter("limit", "", "3")
	i := sub.List(params)
	for i.Next() {
		s := i.Subscription()
		toString, err := json.Marshal(s)

		if !internal.HandleError(err, w) {
			w.Write(toString)
			w.WriteHeader(http.StatusOK)
		}
	}
}
