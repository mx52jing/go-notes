package main

import (
	"bytes"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"go-rabbitmq/shared"
	"time"
)

// worker 下面ch.Consume产生的所有接收器都是哪个先完成就先处理

func createReceiverAndReceiveMessage() {
	// 1、create connection
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	shared.FailOnError(err, "create connection error")
	defer conn.Close()

	// open a channel
	ch, err := conn.Channel()
	shared.FailOnError(err, "open a channel error")
	defer ch.Close()

	// 3、declare queue
	queue, err := ch.QueueDeclare(
		"second_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	shared.FailOnError(err, "declare queue error")
	err = ch.Qos(1, 0, false)
	shared.FailOnError(err, "set prefetch error")

	waitCh := make(chan struct{})
	for i := 1; i <= 3; i++ {
		go consumeMessage(ch, queue.Name, i)
	}
	fmt.Println("[*] Waiting for messages. To exit press CTRL+C")
	<-waitCh
}

func consumeMessage(ch *amqp.Channel, queueName string, serialNumber int) {
	// 4、consume message
	message, err := ch.Consume(
		queueName,
		"",    // consumer
		false, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	shared.FailOnError(err, "consume queue message error")
	go func() {
		for msg := range message {
			dotCount := bytes.Count(msg.Body, []byte("."))
			duration := time.Duration(10 - dotCount)
			fmt.Printf("【NO %d】 Received a message: %s, Will wait [%s]\n", serialNumber, msg.Body, duration*time.Second)
			time.Sleep(duration * time.Second)
			fmt.Printf("【NO %d】Done \n", serialNumber)
			// Manually confirm that the message has been delivered
			msg.Ack(false)
		}
	}()
}

func main() {
	createReceiverAndReceiveMessage()
}
