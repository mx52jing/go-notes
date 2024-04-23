package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

var serverPort = flag.String("port", "8900", "http server port")

/**
大致流程：
1、hub中心保存了所有的client，负责client的注册存储和注销，hub中有一个广播，用来接收某个client传出来的数据，当广播接收到client的数据时，会将数据广播给每个client
2、服务端每次接收到/ws请求，就会创建一个client，然后将client注册到hub中，每个用户都是一个client，client接收用户端发过来的数据，然后将该数据发送给hub中的广播broadcast
*/

func handleMain(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		http.Error(writer, "Not Found", http.StatusNotFound)
		return
	}
	if request.Method != http.MethodGet {
		http.Error(writer, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(writer, request, "index.html")
}

func main() {
	flag.Parse()
	hub := NewHub()
	// 因为Run函数中有for循环，所以需要开协程来运行hub，保证下面的代码可以运行
	go hub.Run()
	server := &http.Server{
		Addr:              fmt.Sprintf(":%s", *serverPort),
		ReadHeaderTimeout: 3 * time.Second,
	}
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		handleClientConnect(writer, request, hub)
	})
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
