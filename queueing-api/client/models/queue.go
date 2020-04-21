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
	URL              string     `json:"url"`
	Name             string     `json:"name"`
	Consumer         string     `json:"consumer"`
	AutoAck          bool       `json:"autoack"`
	Exclusive        bool       `json:"exclusive"`
	NoLocal	         bool       `json:"nolocal"`
	NoWait           bool       `json:"nowait"`
	Arguments        amqp.Table `json:"arguments,omitempty"`
}
