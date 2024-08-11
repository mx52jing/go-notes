package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"go-rabbitmq/shared"
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

	// declare unnamed queue
	queue, err := ch.QueueDeclare("", false, false, true, false, nil)
	shared.FailOnError(err, "declare queue error")

	// bind exchange and queue
	err = ch.QueueBind(queue.Name, "", "logsExchange", false, nil)
	shared.FailOnError(err, "bind exchange and queue error")

	// consume message
	lockChan := make(chan struct{})
	messages, err := ch.Consume(queue.Name, "", true, false, false, false, nil)
	shared.FailOnError(err, "consume message error")
	go func() {
		for msg := range messages {
			fmt.Printf("Received a message: %s\n", msg.Body)
		}
	}()
	fmt.Println("[*] Waiting for messages. To exit press CTRL+C")
	<-lockChan
}

func main() {
	startUpAndReceive()
}
