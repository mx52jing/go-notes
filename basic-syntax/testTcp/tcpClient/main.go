package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	connect, err := net.DialTCP("tcp", nil, &net.TCPAddr{net.ParseIP("127.0.0.1"), 9173, ""})
	defer connect.Close()
	if err != nil {
		fmt.Println("net.DialTCP err", err)
	}
	go func ()  {
		for {
			buf := make([]byte, 1024)
			n, err2 := connect.Read(buf)
			if err2 == io.EOF {
				break
			}
			fmt.Println("客户端接收到的数据为:", string(buf[0:n]))
		}
	}()
	var str string
	fmt.Println("输入文字查看返回吧:")
	for {
		fmt.Scanln(&str)
		if str == "q" {
			break
		}
		// 向服务端发送数据
		connect.Write([]byte (str))
	}
}