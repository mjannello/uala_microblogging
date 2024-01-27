package eventpublisher

import (
	"uala/internal/command/event"
)

type EventPublisher interface {
	Publish(event event.Event) error
}
