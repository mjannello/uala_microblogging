package repository

import "uala/internal/command/event"

type EventStore interface {
	SaveEvent(event event.Event) (int64, error)
}
