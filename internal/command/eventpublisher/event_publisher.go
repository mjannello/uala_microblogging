package eventpublisher

import (
	"uala/internal/model"
)

type EventPublisher interface {
	Publish(event model.Event) error
}
