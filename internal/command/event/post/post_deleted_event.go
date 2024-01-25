package post

import "uala/internal/command/event"

type PostDeletedEvent struct {
	event.Event
}

func (pde *PostDeletedEvent) Type() string {
	return "PostDeletedEvent"
}
