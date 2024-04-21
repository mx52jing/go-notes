package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"os"
)

func main() {
	dialer := websocket.Dialer{}
	// 连接服务端websocket
	conn, _, err := dialer.Dial("ws://192.168.1.2:8900/chat", nil)
	defer conn.Close()
	if err != nil {
		return
	}
	go func() {
		err := sendStdinMessage(conn)
		if err != nil {
			fmt.Printf("client1 send message error: %s\n", err)
		}
	}()
	if err != nil {
		return
	}
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			break
		}
		fmt.Printf("client1 received message is: %s\n", string(p))
	}
	fmt.Printf("client1 closed")
}

func sendStdinMessage(conn *websocket.Conn) error {
	for {
		reader := bufio.NewReader(os.Stdin)
		line, _, err := reader.ReadLine()
		if err != nil {
			return err
		}
		err = conn.WriteMessage(websocket.TextMessage, line)
		if err != nil {
			return err
		}
	}
}
