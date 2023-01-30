package app

import (
	"context"
	"log"
	"rabbitv2/helpers"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Run() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	helpers.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	helpers.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"datastock",
		false,
		false,
		false,
		false,
		nil,
	)
	helpers.FailOnError(err, "Failed to declare a queue")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := "Hello World!"
	err = ch.PublishWithContext(ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	helpers.FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}
