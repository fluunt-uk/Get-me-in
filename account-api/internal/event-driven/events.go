package event_driven

import "github.com/ProjectReferral/Get-me-in/account-api/configs"

var MQ *RabbitClient

func BroadcastUserCreatedEvent(body string) {
	uId := NewUUID()

	//send to fanout exchange
	MQ.Connect()
	MQ.SendToQ(body, configs.FANOUT_EXCHANGE, uId)
}
