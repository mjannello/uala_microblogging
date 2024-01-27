package rabbitmq

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"uala/internal/command/event"
	_ "uala/internal/command/event"
	"uala/internal/command/eventpublisher"
)

const rabbitMQURL = "amqp://guest:guest@rabbitmq:5672/"

type rabbitMQEventPublisher struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

func NewRabbitMQEventPublisher() (eventpublisher.EventPublisher, error) {
	connection, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		return nil, err
	}

	channel, err := connection.Channel()
	if err != nil {
		return nil, err
	}

	err = setupExchangeAndQueues(channel)

	return &rabbitMQEventPublisher{
		Connection: connection,
		Channel:    channel,
	}, nil
}

func setupExchangeAndQueues(channel *amqp.Channel) error {
	err := channel.ExchangeDeclare(
		"event_exchange",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	_, err = channel.QueueDeclare(
		"event_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = channel.QueueBind(
		"event_queue",
		"event_routing_key",
		"event_exchange",
		false,
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}

func (p *rabbitMQEventPublisher) Publish(event event.Event) error {
	eventData, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = p.Channel.Publish(
		"event_exchange",
		"event_routing_key",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        eventData,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
