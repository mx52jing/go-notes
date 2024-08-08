package main

import (
	"github.com/rabbitmq/amqp091-go"
	"go-rabbitmq/shared"
	"log"
)

func createReceiverAndReceiveMessage() {
	// 1、create connection
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	shared.FailOnError(err, "create connection error")
	defer conn.Close()

	// open a channel
	ch, err := conn.Channel()
	shared.FailOnError(err, "open a channel error")
	defer ch.Close()

	// 3、declare queue
	queue, err := ch.QueueDeclare(
		"first_queue",
		false,
		false,
		false,
		false,
		nil,
	)
	shared.FailOnError(err, "declare queue error")
	// 4、consume message
	message, err := ch.Consume(
		queue.Name,
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	shared.FailOnError(err, "consume queue message error")
	waitCh := make(chan struct{})
	go func() {
		for msg := range message {
			log.Printf("Received a message: %s \n", msg.Body)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-waitCh
}

func main() {
	createReceiverAndReceiveMessage()
}
