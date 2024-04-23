package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"math/rand"
	"net/http"
	"time"
)

// 生成指定位数的随机字符串
func RandStringRunes(n int) string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rnd.Intn(len(letterRunes))]
	}
	return string(b)
}

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	WriteBufferSize: 1024,
	ReadBufferSize:  1024,
}

type UserInfo struct {
	UserName string // 用户名
	UserId   string // 用户id
	Message  string // 用户消息
}

type Client struct {
	UserInfo
	hub  *Hub            // hub调度中心
	conn *websocket.Conn // websocket连接
	send chan []byte     // client接受/要发送的数据
}

// 读取到前端发送的消息后 发送给hub
func (client *Client) ReadPump() {
	// ReadPump函数退出的时候关闭连接，并且还要注销client
	defer func() {
		fmt.Printf("【ReadPump】用户【%s:%s】关闭了连接\n", client.UserId, client.UserName)
		client.hub.clientUnRegister <- client // 从hub中注销client
		err := client.conn.Close()
		if err != nil {
			fmt.Printf("【ReadPump】client conn close fail => %s\n", err)
			return
		}
	}()
	for {
		_, message, err := client.conn.ReadMessage()
		if err != nil {
			return
		}
		// 处理以下消息 将换行符\n替换为空格
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		// 默认第一条消息是注册用户
		if len(client.UserId) == 0 {
			userName := string(message)
			userId := RandStringRunes(16)
			client.UserName = userName
			client.UserId = userId
			userBytesData, err := json.Marshal(UserInfo{
				UserId: userId, UserName: userName, Message: "",
			})
			if err != nil {
				return
			}
			client.hub.broadcast <- userBytesData
		} else {
			userSendData, err := json.Marshal(UserInfo{
				UserId: client.UserId, UserName: client.UserName, Message: string(message),
			})
			if err != nil {
				return
			}
			// 将消息传给广播
			client.hub.broadcast <- userSendData
		}
	}
}

// 接收到hub的广播消息 发送给前端
func (client *Client) WritePump() {
	defer func() {
		fmt.Printf("【WritePump】用户【%s:%s】关闭了连接\n", client.UserId, client.UserName)
		err := client.conn.Close()
		if err != nil {
			fmt.Printf("【WritePump】client conn close fail => %s\n", err)
			return
		}
	}()
	for {
		message, ok := <-client.send
		// ok为false的时候 说明管道被close了
		if !ok {
			return
		}
		//client.conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("buy buy %s", client.UserName)))
		err := client.conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			fmt.Printf("向浏览器发送数据失败: %v\n", err)
			return
		}
		return
	}
}

func handleClientConnect(writer http.ResponseWriter, request *http.Request, hub *Hub) {
	// 升级http为websocket
	conn, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		fmt.Printf("updrade http to websocket fail => %s\n", err)
		return
	}
	// 实例化一个client
	client := &Client{
		hub:  hub,
		send: make(chan []byte),
		conn: conn,
	}
	// 注册client
	client.hub.clientRegister <- client
	// 监听接收数据/发送数据
	go client.ReadPump()
	go client.WritePump()
}
