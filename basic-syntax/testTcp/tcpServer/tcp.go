package main

import (
	"fmt"
	"io"
	"net"
)

func testTcp() {
	// 创建tcp监听
	listen, listenErr := net.ListenTCP("tcp", &net.TCPAddr{net.ParseIP("127.0.0.1"), 9173, ""})
	if listenErr != nil {
		fmt.Println("net.ListenTCP err", listenErr)
		return
	}
	fmt.Println("等待连接......")
	// 等待连接
	for {
		connect, err := listen.AcceptTCP()
		defer connect.Close()
		if err != nil {
			fmt.Println("listen.AcceptTCP() err", err)
			break;
		}
		// 获取客户端连接地址
		addr := connect.RemoteAddr().String()
		fmt.Println("客户端地址为：", addr)
		for {
			buf := make([]byte, 1024)
			n, err2 := connect.Read(buf)
			if err2 != nil {
				if err2 != io.EOF {
					fmt.Println("Error reading from client:", err)
				}
				break
			}
			fmt.Println("从客户端获取到的数据为:", string(buf[0:n]))
			// 发送响应数据给客户端
			connect.Write([]byte (string(buf[0:n])))
		}
	}
}


func main() {
	testTcp()
}