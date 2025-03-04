package repositories

import (
	"consumer/src/api/config"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func ConnectRMQ() *amqp.Channel {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s", config.RabbitMQUser, config.RabbitMQPassword, config.RabbitMQHostname, config.RabbitMQPort))
	failOnError(err, "Failed to connect to RabbitMQ")
	// defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	// defer ch.Close()

	DeclareQueue(ch)
	return ch
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
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
	failOnError(err, "Failed to declare a queue")
}
