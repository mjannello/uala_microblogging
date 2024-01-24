package post

import "uala/internal/command/event"

type PostDeletedEvent struct {
	event.EventBase
	ID      string
	Content string
}

func (pde *PostDeletedEvent) Type() string {
	return "PostDeletedEvent"
}
