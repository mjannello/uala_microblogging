package post

import "uala/internal/command/event"

type PostAddedEvent struct {
	event.Event
}

func (pae *PostAddedEvent) Type() string {
	return "PostAddedEvent"
}
