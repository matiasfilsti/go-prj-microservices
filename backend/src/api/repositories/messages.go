package repositories

import (
	errors "backend/src/api/domain/errors"
	"context"

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

func (m *MessageRepository) Send(ctx context.Context) error {
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
	if err != nil {
		return errors.NewRabbitmqMsgError("error sending message to rabbit")
	}

}
