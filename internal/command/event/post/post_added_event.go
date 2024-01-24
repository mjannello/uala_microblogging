package post

import "uala/internal/command/event"

type PostAddedEvent struct {
	event.EventBase
	ID      string
	Content string
}

func (pae *PostAddedEvent) Type() string {
	return "PostAddedEvent"
}
