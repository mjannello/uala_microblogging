package rabbitmq

import (
	"uala/internal/command/eventpublisher"
)

type rabbitMQEventPublisher struct {
	// TODO: Set RabbitMQ config
}

func NewRabbitMQEventPublisher() eventpublisher.EventPublisher {
	return &rabbitMQEventPublisher{}
}

func (rep *rabbitMQEventPublisher) Publish(topic string, eventData interface{}) error {
	// TODO: Implement RabbitMQ publish method
	return nil
}
