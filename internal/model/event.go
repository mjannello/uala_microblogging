package model

import "time"

type Event interface {
	GetType() string
	GetContent() string
	GetUserName() string
	GetDate() time.Time
	GetData() map[string]interface{}
}
