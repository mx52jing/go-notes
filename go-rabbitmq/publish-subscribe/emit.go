package main

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"go-rabbitmq/shared"
	"strings"
	"time"
)

func startUpAndSend() {
	// create connection
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	shared.FailOnError(err, "create connection error")
	defer conn.Close()

	// create channel
	ch, err := conn.Channel()
	shared.FailOnError(err, "create channel error")
	defer ch.Close()

	// declare exchange
	err = ch.ExchangeDeclare(
		"logsExchange",
		"fanout", // exchange type => headers, topic, direct, fanout
		false,    // durable
		false,    // auto-delete
		false,    // internal
		false,    // no-wait
		nil,      // args
	)
	shared.FailOnError(err, "declare exchange error")

	// publish message
	ctx, cancelFunc := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancelFunc()
	for i := 0; i < 6; i++ {
		go func(i int) {
			messageBody := fmt.Sprintf("我是第[%d]条消息%s", i+1, strings.Repeat(".", i+1))
			err = ch.PublishWithContext(
				ctx,
				"logsExchange",
				"",
				false,
				false,
				amqp.Publishing{
					DeliveryMode: amqp.Persistent,
					ContentType:  "text/plain",
					Body:         []byte(messageBody),
				},
			)
			shared.FailOnError(err, "publish message error")
			fmt.Printf(" [x] Sent %s\n", messageBody)
		}(i)
	}
	time.Sleep(6 * time.Second)
}

func main() {
	startUpAndSend()
}
