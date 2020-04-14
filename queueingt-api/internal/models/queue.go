package models

import "github.com/streadway/amqp"

//'json:' is the value that will be picked up from the JSON body
//JSON must contain the value after 'json:...'  instead of the attribute name

type QueueDeclare struct {
	Name             string     `json:"name"`
	Durable          bool       `json:"durable"`
	DeleteWhenUnused bool       `json:"deletewhenunused"`
	Exclusive        bool       `json:"exclusive"`
	NoWait           bool       `json:"nowait"`
	Arguments        amqp.Table // arguments unsuported
}

type ExchangeDeclare struct {
	Name             string     `json:"name"`
	Type             string     `json:"type"`
	Durable          bool       `json:"durable"`
	DeleteWhenUnused bool       `json:"deletewhenunused"`
	Exclusive        bool       `json:"exclusive"`
	NoWait           bool       `json:"nowait"`
	Arguments        amqp.Table // arguments unsuported
}

type QueueBind struct {
	Name             string     `json:"name"`
	Key              string     `json:"key"`
	Exchange         bool       `json:"exchange"`
	DeleteWhenUnused bool       `json:"deletewhenunused"`
	Arguments        amqp.Table // arguments unsuported
}

type ExchangePublish struct {
	Exchange        string           `json:"exchange"`
	key             string           `json:"key"`
	Mandatory       bool             `json:"mandatory"`
	Immediate       bool             `json:"immediate"`
	Publishing       amqp.Publishing `json:"publishing"`
}