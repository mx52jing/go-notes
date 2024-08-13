package main

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"go-rabbitmq/shared"
	"os"
	"time"
)

func startUpAndEmit() {
	// create connection
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	shared.FailOnError(err, "create connection error")
	defer conn.Close()

	// create channel
	ch, err := conn.Channel()
	shared.FailOnError(err, "create channel error")
	defer ch.Close()
	exchangeName := "topicLogs"
	// declare exchange
	err = ch.ExchangeDeclare(
		exchangeName,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	shared.FailOnError(err, "declare exchange error")

	// publish message
	ctx, cancelFunc := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancelFunc()
	message := shared.BodyFrom(os.Args)

	err = ch.PublishWithContext(
		ctx,
		exchangeName,
		shared.GetLogSeverity(os.Args),
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	shared.FailOnError(err, "publish message error")
	fmt.Printf("Successful Send message [%s]\n", message)
}

func main() {
	startUpAndEmit()
}
