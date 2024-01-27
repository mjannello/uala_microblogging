package event

import "time"

type Event interface {
	GetType() string
	GetContent() string
	GetUserName() string
	GetDate() time.Time
}
