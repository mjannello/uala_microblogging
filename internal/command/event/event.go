package event

import "time"

type Event interface {
	Type()
	GetDateCreated()
}

type EventBase struct {
	DateCreated time.Time
}

func (eb *EventBase) GetDateCreated() time.Time {
	return eb.DateCreated
}
