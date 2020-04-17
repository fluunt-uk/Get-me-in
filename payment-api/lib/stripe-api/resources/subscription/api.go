package subscription

import (
	"fmt"
	stripe_api "github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api"
	"github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api/resources/models"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/sub"
	"net/http"
)

func CreateSub(w http.ResponseWriter, r *http.Request) {
	params := &stripe.SubscriptionParams{
		Customer: stripe.String("cus_H7HyJY5cWLA7Uf"),
		Items: []*stripe.SubscriptionItemsParams{
			{
				Plan: stripe.String("plan_H4eVnOxhxYYZ7a"),
			},
		},
	}
	s, _ := sub.New(params)

	stripe_api.ReturnSuccessJSON(w, &s)

	status, err := AddSubscription(models.Subscription{
		Email:          "hamza@gmail.com",
		AccountID:      s.Customer.ID,
		SubscriptionID: s.ID,
		PlanID:         s.Plan.ID,
		PlanType:       "Hamuzzz",
	})

	if err != nil{
		fmt.Println(status)
	}
	fmt.Println(status, err)
}

func RetrieveSub(w http.ResponseWriter, r *http.Request) {
	s, _ := sub.Get("sub_H6qCxUjOuCCmfj", nil)

	stripe_api.ReturnSuccessJSON(w, &s)
}

func UpdateSub(w http.ResponseWriter, r *http.Request) {
	params := &stripe.SubscriptionParams{}
	params.AddMetadata("order_id", "0001")
	s, _ := sub.Update("sub_H6qCxUjOuCCmfj", params)

	stripe_api.ReturnSuccessJSON(w, &s)
}

func CancelSub(w http.ResponseWriter, r *http.Request) {
	s, _ := sub.Cancel("sub_H6qCxUjOuCCmfj", nil)

	stripe_api.ReturnSuccessJSON(w, &s)
	//status, err = DeleteSubscription()
}

func ListSubs(w http.ResponseWriter, r *http.Request) {
	params := &stripe.SubscriptionListParams{}
	params.Filters.AddFilter("limit", "", "3")
	i := sub.List(params)
	for i.Next() {
		s := i.Subscription()

		stripe_api.ReturnSuccessJSON(w, &s)
	}
}
