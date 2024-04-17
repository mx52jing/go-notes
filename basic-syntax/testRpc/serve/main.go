package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

type Server struct {}

func (s *Server) Sum(args *Args, reply *int) error {
	*reply = args.M + args.N
	return nil
}

type Args struct {
	M, N int
}

// 创建一个用于计算的对象Server，并且将它通过 rpc.Register(object) 注册，
// 调用 rpc.HandleHTTP方法
// 调用`net.Listen`方法监听
// 调用`http.Serve`启动服务

func createServer() {
	calc := new(Server)
	rpc.Register(calc)
	rpc.HandleHTTP()
	listen, err := net.Listen("tcp", ":9100")
	if err != nil {
		fmt.Println("listen err", err)
		return
	}
	http.Serve(listen, nil)
}


func main() {
	createServer()
}