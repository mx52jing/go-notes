package main

import (
	"fmt"
	"net/rpc"
)

type Args struct {
	M, N int
}

func main() {
	client, err := rpc.DialHTTP("tcp", ":9100")
	if err != nil {
		fmt.Println("rpc.DialHTTP err", err)
		return
	}
	args := Args {10, 20}
	var res int
	client.Call("Server.Sum", &args, &res)
	fmt.Println(res, "计算后的值")
}