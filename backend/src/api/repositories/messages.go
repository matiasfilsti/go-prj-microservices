package repositories

import (
	"context"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type MessageRepository struct {
	ch *amqp.Channel
}

func NewMessageRepository(ch *amqp.Channel) MessageRepository {
	return MessageRepository{
		ch: ch,
	}

}

func (m *MessageRepository) Send(ctx context.Context) {
	body := "Hello World!"
	err := m.ch.PublishWithContext(ctx,
		"",      // exchange
		"hello", // routing key
		false,   // mandatory
		false,   // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)

}
