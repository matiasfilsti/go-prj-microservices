package repositories

import (
	"backend/src/api/config"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func ConnectRMQ() *amqp.Channel {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s", config.RabbitMQUser, config.RabbitMQPassword, config.RabbitMQHostname, config.RabbitMQPort))
	if err != nil {
		log.Printf("Warning: Failed to connect to RabbitMQ: %v", err)
		return nil
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Printf("Warning: Failed to open RabbitMQ channel: %v", err)
		conn.Close()
		return nil
	}

	DeclareQueue(ch)
	return ch
}

func DeclareQueue(ch *amqp.Channel) {
	_, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Printf("Warning: Failed to declare queue: %v", err)
		return
	}
}
