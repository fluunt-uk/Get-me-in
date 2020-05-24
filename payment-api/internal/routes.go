package internal

import (
	"github.com/ProjectReferral/Get-me-in/payment-api/configs"
	"github.com/ProjectReferral/Get-me-in/payment-api/internal/service"
	"github.com/ProjectReferral/Get-me-in/pkg/security"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)


type EndpointBuilder struct {
	router       	*mux.Router
	ss 				service.Subscription
}

func (eb *EndpointBuilder) SetupRouter(route *mux.Router) {
	eb.router = route
}

func (eb *EndpointBuilder) InjectSubscriptionServ(ss service.Subscription) {
	eb.ss = ss
}

func (eb *EndpointBuilder) SetupEndpoints() {

	eb.router.HandleFunc("/premium/subscribe", security.WrapHandlerWithSpecialAuth(eb.ss.SubscribeToPremiumPlan, configs.AUTH_AUTHENTICATED)).Methods("POST")
	log.Fatal(http.ListenAndServe(configs.PORT, eb.router))
}
