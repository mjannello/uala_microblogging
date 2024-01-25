package post

import "uala/internal/command/event"

type PostUpdatedEvent struct {
	event.Event
}

func (pue *PostUpdatedEvent) Type() string {
	return "PostUpdatedEvent"
}
