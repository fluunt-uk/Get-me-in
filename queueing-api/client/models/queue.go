package models

import (
	"github.com/streadway/amqp"
	"time"
)

//'json:' is the value that will be picked up from the JSON body
//JSON must contain the value after 'json:...'  instead of the attribute name

type SubscribeMessage interface {
	GetID() string
}

type QueueDeclare struct {
	Name             string     `json:"name"`
	Durable          bool       `json:"durable"`
	DeleteWhenUnused bool       `json:"deletewhenunused"`
	Exclusive        bool       `json:"exclusive"`
	NoWait           bool       `json:"nowait"`
	Arguments        amqp.Table `json:"arguments,omitempty"`
}

type ExchangeDeclare struct {
	Name             string     `json:"name"`
	Type             string     `json:"type"`
	Durable          bool       `json:"durable"`
	DeleteWhenUnused bool       `json:"deletewhenunused"`
	Exclusive        bool       `json:"exclusive"`
	NoWait           bool       `json:"nowait"`
	Arguments        amqp.Table `json:"arguments,omitempty"`
}

type QueueBind struct {
	Name             string     `json:"name"`
	Key              string     `json:"key"`
	Exchange         string     `json:"exchange"`
	DeleteWhenUnused bool       `json:"deletewhenunused"`
	Arguments        amqp.Table `json:"arguments,omitempty"`
}

type ExchangePublish struct {
	Exchange        string           `json:"exchange"`
	Key             string           `json:"key"`
	Mandatory       bool             `json:"mandatory"`
	Immediate       bool             `json:"immediate"`
	Publishing      amqp.Publishing  `json:"publishing"`
}

type QueueConsume struct {
	Name             string     `json:"name"`
	Consumer         string     `json:"consumer"`
	AutoAck          bool       `json:"autoack"`
	Exclusive        bool       `json:"exclusive"`
	NoLocal	         bool       `json:"nolocal"`
	NoWait           bool       `json:"nowait"`
	Arguments        amqp.Table `json:"arguments,omitempty"`
}

//configure how often messages will be delivered
type QueueQos struct {
	PrefetchCount    int        `json:"prefetchcount"`
	PrefetchSize     int        `json:"prefetchsize"`
}

type QueueSubscribe struct {
	URL              string          `json:"url"`
	Name             string          `json:"name"`
	Consumer         string          `json:"consumer"`
	Exclusive        bool            `json:"exclusive"`
	NoLocal	         bool            `json:"nolocal"`
	NoWait           bool            `json:"nowait"`
	MaxRetry         int             `json:"maxretry"`
	Timeout          time.Duration   `json:"timeout"`
	Qos              QueueQos        `json:"qos"`
	Arguments        amqp.Table      `json:"arguments,omitempty"`
}

type QueueSubscribeId struct {
	ID           string     `json:"id"`
}

type QueueMessage struct {
	ID               uint64     `json:"id"`
	RetryCount       uint64     `json:"retrycount"`
	Body           []byte       `json:"body"`
}

type QueueFailedMessage struct {
	ID               uint64     `json:"id"`
	Body           []byte       `json:"body"`
	RetryCount       uint64     `json:"retrycount"`
	Reason           string     `json:"reason"`
}

type MessageAcknowledge struct {
	SubID            QueueSubscribeId  `json:"subID"`
	ID               uint64            `json:"id"`
	Body           []byte              `json:"body"`	
	Acknowledge      bool              `json:"acknowledge"`
	Requeue          bool              `json:"requeue,omitempty"` //only valid if acknowledge is false
	Multiple         bool              `json:"multiple"`
}

type MessageReject struct {
	SubID            QueueSubscribeId  `json:"subID"`
	ID               uint64            `json:"id"`
	Body           []byte              `json:"body"`
	Requeue          bool              `json:"requeue"`
}

func (ma MessageAcknowledge) GetID() string {
	return ma.SubID.ID
}

func (mr MessageReject) GetID() string {
	return mr.SubID.ID
}