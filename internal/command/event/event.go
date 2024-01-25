package event

import "time"

type Event struct {
	ID          uint64
	Type        string
	Content     string
	DateCreated time.Time
}

func (e *Event) GetDateCreated() time.Time {
	return e.DateCreated
}
