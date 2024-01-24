package rabbitmq

type RabbitMQEventPublisher struct {
	// TODO: Set RabbitMQ config
}

func (rep *RabbitMQEventPublisher) Publish(topic string, eventData interface{}) error {
	// TODO: Implement RabbitMQ publish method
	return nil
}
