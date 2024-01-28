package rabbitmq_consumer

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"uala/internal/query/eventconsumer"
	"uala/internal/query/service"
	"uala/pkg/logger"
)

const rabbitMQURL = "amqp://guest:guest@rabbitmq:5672/"

type rabbitMQEventConsumer struct {
	queryService service.QueryService
	conn         *amqp.Connection
	channel      *amqp.Channel
}

func NewRabbitMQEventConsumer(s service.QueryService) (eventconsumer.EventConsumer, error) {
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &rabbitMQEventConsumer{
		queryService: s,
		conn:         conn,
		channel:      channel,
	}, nil
}

func (c *rabbitMQEventConsumer) StartConsuming() error {
	defer func() {
		_ = c.conn.Close()
		_ = c.channel.Close()
	}()

	msgs, err := c.channel.Consume(
		"event_queue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	for msg := range msgs {
		if err := c.consumeEvent(msg.Body); err != nil {
			log.Printf("Error consuming event: %v", err)
		}
	}

	return nil
}

func (c *rabbitMQEventConsumer) consumeEvent(data []byte) error {
	var eventData map[string]interface{}
	if err := json.Unmarshal(data, &eventData); err != nil {
		return err
	}
	logger.Logger.Print("consume event eventData", eventData)
	if err := c.queryService.UpdateRepositoryWithEvent(eventData); err != nil {
		log.Printf("Error updating repository with event: %v", err)
		return err
	}

	return nil
}
