package eventpublisher

type EventPublisher interface {
	Publish(topic string, eventData interface{}) error
}
