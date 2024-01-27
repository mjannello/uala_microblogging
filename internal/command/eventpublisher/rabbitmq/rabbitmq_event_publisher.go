package rabbitmq

import (
	"uala/internal/command/event"
	"uala/internal/command/eventpublisher"
)

type rabbitMQEventPublisher struct {
	// TODO: Set RabbitMQ config
}

func NewRabbitMQEventPublisher() eventpublisher.EventPublisher {
	return &rabbitMQEventPublisher{}
}

func (rep *rabbitMQEventPublisher) Publish(event event.Event) error {
	// TODO: Implement RabbitMQ publish method
	return nil
}
