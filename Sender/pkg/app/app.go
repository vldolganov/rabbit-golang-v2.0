package app

import (
	"context"
	"log"
	"rabbitv2/pkg/helpers"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Run() {

	companies := [5]string{
		"IBM",
		"AAPL",
		"INTC",
		"AMD",
		"TSLA",
	}

	conn, err := amqp.Dial("amqp://admin:root@localhost:5672/")
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

	for range time.Tick(2 * time.Second) {
		for i := range companies {
			body := helpers.DataParser(companies[i])
			err = ch.PublishWithContext(ctx,
				"",
				q.Name,
				false,
				false,
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        body,
				})
			helpers.FailOnError(err, "Failed to publish a message")
			log.Printf(" [x] Sent %s\n", body)
		}
	}
}
