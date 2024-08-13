package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"go-rabbitmq/shared"
	"os"
)

func startUpAndReceive() {
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

	// declare queue
	queue, err := ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	shared.FailOnError(err, "declare queue error")

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [binding_key]...", os.Args[0])
		os.Exit(1)
	}
	// bind queue and exchange
	for _, routingKey := range os.Args[1:] {
		fmt.Printf("Binding queue [%s] to exchange [%s] with routing key [%s]\n", queue.Name, "topicLogs", routingKey)
		err = ch.QueueBind(
			queue.Name,
			routingKey,
			"topicLogs",
			false,
			nil,
		)
		shared.FailOnError(err, "bind queue error")
	}

	// consume message
	messages, err := ch.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	shared.FailOnError(err, "consume message error")

	lockChan := make(chan struct{})
	go func() {
		for delivery := range messages {
			fmt.Printf("Receive message %s\n", delivery.Body)
		}
	}()
	<-lockChan
}

func main() {
	startUpAndReceive()
}
