package event_driven

import "github.com/ProjectReferral/Get-me-in/account-api/configs"

func BroadcastUserCreatedEvent(body string){
	uId := NewUUID()

	//send to fanout exchange
	SendToQ(body, configs.FANOUT_EXCHANGE, uId)
}
