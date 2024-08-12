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

	// publish message
	ctx, cancelFunc := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancelFunc()

	// 获取运行文件时输入的文案，将该文案作为发送的message
	message := shared.BodyFrom(os.Args)
	err = ch.PublishWithContext(
		ctx,
		"directLogs",
		shared.GetLogSeverity(os.Args), // routing key 获取命令行发送的日志等级
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	shared.FailOnError(err, "publish message error")
	fmt.Printf("send message %s\n", message)
}

func main() {
	startUpAndEmit()
}
