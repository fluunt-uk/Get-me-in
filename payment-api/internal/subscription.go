package internal

import (
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/sub"
)

func createSub() (*stripe.Subscription, error) {
	params := &stripe.SubscriptionParams{
		Customer: stripe.String("cus_H4HfdRtWmkH717"),
		Items: []*stripe.SubscriptionItemsParams{
			{
				Plan: stripe.String("plan_EeE4ns3bvb34ZP"),
			},
		},
	}
	s, err := sub.New(params)

	return s, err
}

func retrieveSub() *stripe.Subscription {
	s, _ := sub.Get("sub_36VrPHS2vVxJMq", nil)
	return s
}

func updateSub() *stripe.Subscription {
	params := &stripe.SubscriptionParams{}
	params.AddMetadata("order_id", "6735")
	s, _ := sub.Update("sub_36VrPHS2vVxJMq", params)

	return s
}

func cancelSub() *stripe.Subscription {
	s, _ := sub.Cancel("sub_36VrPHS2vVxJMq", nil)

	return s
}

func listSubs() *stripe.Subscription {
	params := &stripe.SubscriptionListParams{}
	params.Filters.AddFilter("limit", "", "3")
	i := sub.List(params)
	for i.Next() {
		s := i.Subscription()
		return s
	}
	return nil
}
