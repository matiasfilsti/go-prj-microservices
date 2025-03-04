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

func (m *MessageRepository) Read() {
	msgs, err := m.ch.Consume(
		"hello",      // queue
		"consumer-1", // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}
