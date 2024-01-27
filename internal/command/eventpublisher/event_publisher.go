package eventpublisher

import (
	"fmt"
	"uala/internal/command/event"
)

type EventPublisher interface {
	Publish(event event.Event) error
}

type eventPublisher struct {
	Topic string
}

func NewEventPublisher(topic string) EventPublisher {
	return &eventPublisher{
		Topic: topic,
	}
}

func (ep *eventPublisher) Publish(event event.Event) error {
	return fmt.Errorf("must be implemented in specific publisher")
}
