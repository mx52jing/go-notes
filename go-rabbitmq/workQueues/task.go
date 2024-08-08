package main

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"go-rabbitmq/shared"
	"strings"
	"time"
)

func connectAndSendMessage() {
	// 1. connect to the RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	shared.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// 2. open a channel
	ch, err := conn.Channel()
	shared.FailOnError(err, "create channel error")
	defer ch.Close()

	// 3. declare a queue，it will only be created if it doesn't exist already
	queue, err := ch.QueueDeclare(
		"second_queue", // queue name
		true,           // durable
		false,          //delete when unused
		false,          //exclusive
		false,          // no-wait
		nil,            // arguments
	)
	shared.FailOnError(err, "declare queue error")

	withTimeoutCtx, cancelFunc := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancelFunc()
	// declare a message to be sent

	for i := 0; i < 10; i++ {
		messageBody := fmt.Sprintf("我是第[%d]条消息%s", i+1, strings.Repeat(".", i+1))
		err = ch.PublishWithContext(
			withTimeoutCtx,
			"",         //exchange name
			queue.Name, // routing key
			false,      // mandatory
			false,      //immediate
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				Body:         []byte(messageBody),
				ContentType:  "text/plain",
			},
		)
		shared.FailOnError(err, "Publish message error")
		fmt.Printf(" [x] Sent %s\n", messageBody)
	}
	time.Sleep(6 * time.Second)
}

func main() {
	connectAndSendMessage()
}
