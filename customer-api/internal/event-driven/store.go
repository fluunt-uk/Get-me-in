package event_driven

import "github.com/ProjectReferral/Get-me-in/queueing-api/client/models"

//we store all the new subscribers in this map
type SubscriberStore struct {
	store  map[string]*models.QueueSubscribeId
}

func (ss *SubscriberStore) Init(){
	ss.store = make(map[string]*models.QueueSubscribeId)
}

func (ss *SubscriberStore) AppendSubscriber(id string, sub *models.QueueSubscribeId){
	ss.store[id] = sub
}

func (ss *SubscriberStore) GetSubscriber(id string) *models.QueueSubscribeId{
	return ss.store[id]
}

func (ss *SubscriberStore) RemoveSubscriber(id string){
	delete(ss.store, id)
}

