package card

import (
	"fmt"
	sub_builder "github.com/ProjectReferral/Get-me-in/payment-api/lib/dynamodb/repo"
	stripe_api "github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api"
	"github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api/resources/models"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/sub"
	"net/http"
)

//interface with the implemented methods will be injected in this variable
type Builder interface {
	Put(c *stripe.Customer, pt string) (*models.Subscription, error)
	Get(http.ResponseWriter, *http.Request)
	Cancel(http.ResponseWriter, *http.Request)
	Patch(http.ResponseWriter, *http.Request)
	GetBatch(http.ResponseWriter, *http.Request)
	TestCreate(http.ResponseWriter, *http.Request)
}

type Wrapper struct{
	DynamoSubRepo sub_builder.Builder
}

func (cw *Wrapper) Put(c *stripe.Customer, pt string) (*models.Subscription, error){
	params := &stripe.SubscriptionParams{
		Customer: stripe.String(c.ID),

		Items: []*stripe.SubscriptionItemsParams{
			{
				Plan: stripe.String(pt),
			},
		},
	}
	s, e := sub.New(params)

	if e != nil {
		return nil, e
	}

	var sm = &models.Subscription{
		Email:          c.Email,
		AccountID:      s.Customer.ID,
		SubscriptionID: s.ID,
		PlanID:         s.Plan.ID,
		PlanType:       s.Plan.Nickname,
		Price:			s.Plan.Amount,
	}

	//TODO: might not be needed?
	//status, err := cw.DynamoSubRepo.Create(sm)
	//
	//if err != nil{
	//	fmt.Println(status, err)
	//}

	return sm, nil
}

func (cw *Wrapper) Get(w http.ResponseWriter, r *http.Request) {
	s, _ := sub.Get("sub_H6qCxUjOuCCmfj", nil)

	stripe_api.ReturnSuccessJSON(w, &s)
}

func (cw *Wrapper) Patch(w http.ResponseWriter, r *http.Request) {
	params := &stripe.SubscriptionParams{}
	params.AddMetadata("order_id", "0001")
	s, _ := sub.Update("sub_H6qCxUjOuCCmfj", params)

	stripe_api.ReturnSuccessJSON(w, &s)
}

func (cw *Wrapper) Cancel(w http.ResponseWriter, r *http.Request) {
	s, _ := sub.Cancel("sub_H6qCxUjOuCCmfj", nil)

	stripe_api.ReturnSuccessJSON(w, &s)
	//status, err = DeleteSubscription()
}

//it return 3 ReturnSuccessJSON as per the limit
//but SOMEHOW (to-be figured out) the method is auto called as many times as needed to get all Subs
func (cw *Wrapper) GetBatch(w http.ResponseWriter, r *http.Request) {
	params := &stripe.SubscriptionListParams{}
	//A limit on the number of objects to be returned. Limit can range between 1 and 100, and the default is 10.
	params.Filters.AddFilter("limit", "", "3")
	i := sub.List(params)
	for i.Next() {
		s := i.Subscription()
		stripe_api.ReturnSuccessJSON(w, &s)
	}
}

func (cw *Wrapper) TestCreate(w http.ResponseWriter, r *http.Request) {

	status, err := cw.DynamoSubRepo.Create(&models.Subscription{
		Email:          "hamza@gmail.com",
	})

	if err != nil{
		fmt.Println(status, err)
	}
	fmt.Println(status, err)

}
