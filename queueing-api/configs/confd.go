package configs

import "time"

const (
	PORT = ":5004"
	Threads = 5
	FailedMessageQueue = "failed-message-queue"
	SleepTime = 1 * time.Millisecond
)

var BrokerUrl = ""