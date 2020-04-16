package subscription

import (
	"fmt"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/sub"
	"net/http"
	"github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api/resources/models"
)

func CreateSub(w http.ResponseWriter, r *http.Request) {
	params := &stripe.SubscriptionParams{
		Customer: stripe.String("cus_H6Sh6OUR88nUKr"),
		Items: []*stripe.SubscriptionItemsParams{
			{
				Plan: stripe.String("plan_H4eVnOxhxYYZ7a"),
			},
		},
	}
	s, _ := sub.New(params)

	ReturnSuccessJSON(w, s)

	status, err := AddSubscription(models.Subscription{
		Email:          s.Customer.Email,
		AccountID:      s.Customer.ID,
		SubscriptionID: s.ID,
		PlanID:         s.Plan.ID,
		PlanType:       s.Plan.Nickname,
	})

	if err != nil{
		fmt.Println(status)
	}
	fmt.Println(status, err)
}

func RetrieveSub(w http.ResponseWriter, r *http.Request) {
	s, _ := sub.Get("sub_36VrPHS2vVxJMq", nil)

	ReturnSuccessJSON(w, s)
}

func UpdateSub(w http.ResponseWriter, r *http.Request) {
	params := &stripe.SubscriptionParams{}
	params.AddMetadata("order_id", "6735")
	s, _ := sub.Update("sub_36VrPHS2vVxJMq", params)

	ReturnSuccessJSON(w, s)
}

func CancelSub(w http.ResponseWriter, r *http.Request) {
	s, _ := sub.Cancel("sub_36VrPHS2vVxJMq", nil)

	ReturnSuccessJSON(w, s)

	//status, err = DeleteSubscription()
}

func ListSubs(w http.ResponseWriter, r *http.Request) {
	params := &stripe.SubscriptionListParams{}
	params.Filters.AddFilter("limit", "", "3")
	i := sub.List(params)
	for i.Next() {
		s := i.Subscription()

		ReturnSuccessJSON(w, s)
	}
}
