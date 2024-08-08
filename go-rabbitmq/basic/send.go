package main

import (
	"context"
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"go-rabbitmq/shared"
	"time"
)

func connectAndSendMessage() {
	// 1. connect to the RabbitMQ
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	shared.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// 2. open a channel
	ch, err := conn.Channel()
	shared.FailOnError(err, "create channel error")
	defer ch.Close()

	// 3. declare a queue，it will only be created if it doesn't exist already
	queue, err := ch.QueueDeclare(
		"first_queue", // queue name
		false,         // durable
		false,         //delete when unused
		false,         //exclusive
		false,         // no-wait
		nil,           // arguments
	)
	shared.FailOnError(err, "declare queue error")

	withTimeoutCtx, cancelFunc := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancelFunc()
	// declare a message to be sent
	messageBody := "hello, I am the first message, you have received me"
	err = ch.PublishWithContext(
		withTimeoutCtx,
		"",         //exchange name
		queue.Name, // routing key
		false,      // mandatory
		false,      //immediate
		amqp091.Publishing{
			Body:        []byte(messageBody),
			ContentType: "text/plain",
		},
	)
	shared.FailOnError(err, "Publish error")
	fmt.Printf(" [x] Sent %s\n", messageBody)
}

func main() {
	connectAndSendMessage()
}
