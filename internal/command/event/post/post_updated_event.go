package post

import "uala/internal/command/event"

type PostUpdatedEvent struct {
	event.EventBase
	ID      string
	Content string
}

func (pue *PostUpdatedEvent) Type() string {
	return "PostUpdatedEvent"
}
