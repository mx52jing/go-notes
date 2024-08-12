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

	// declare exchange
	err = ch.ExchangeDeclare(
		"directLogs",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	shared.FailOnError(err, "declare exchange error")

	// queue declare
	queue, err := ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	shared.FailOnError(err, "declare queue error")
	// 绑定多个类型的log
	for _, routingKey := range os.Args[1:] {
		fmt.Printf("Binding queue [%s] to exchange [%s] with routing key [%s]\n", queue.Name, "directLogs", routingKey)
		err = ch.QueueBind(
			queue.Name,
			routingKey,
			"directLogs",
			false,
			nil,
		)
		shared.FailOnError(err, "QueueBind error")
	}

	messages, err := ch.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	shared.FailOnError(err, "consume error")
	lockCh := make(chan struct{})
	go func() {
		for msg := range messages {
			fmt.Printf("接收到的消息为：%s\n", msg.Body)
		}
	}()
	fmt.Println("[*] Waiting for messages. To exit press CTRL+C")
	<-lockCh
}

func main() {
	startUpAndReceive()
}
