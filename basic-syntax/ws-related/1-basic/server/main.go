package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

// 升级到websocket协议的配置
var upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// 所有的连接者
var conns []*websocket.Conn

func handleWs(res http.ResponseWriter, req *http.Request) {
	// 升级到websocket协议
	conn, err := upgrade.Upgrade(res, req, nil)
	conns = append(conns, conn)
	defer conn.Close()
	if err != nil {
		return
	}
	// 要让每个连接都能接受到其他的client发送的消息的话 就要将每个conn存起来
	for {
		// 读取接受到的消息
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("err is %s\n", err)
			break
		}
		for _, curConn := range conns {
			// 将接收到的消息发送回去 发送给每一个client
			sendMeg := fmt.Sprintf("你说的是：%s吗\n", string(message))
			err = curConn.WriteMessage(messageType, []byte(sendMeg))
			if err != nil {
				break
			}
		}
		fmt.Printf("server端收到的messageType is %d，message is %s\n", messageType, string(message))
	}
	fmt.Printf("server端服务关闭")
}

func main() {
	http.HandleFunc("/chat", handleWs)
	err := http.ListenAndServe(":8900", nil)
	if err != nil {
		return
	}
}
