package eventstore

import (
	"uala/internal/model"
)

type EventStore interface {
	SaveEvent(event model.Event) (int64, error)
}
