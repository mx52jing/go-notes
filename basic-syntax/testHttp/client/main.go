package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	res, err := http.Get("http://127.0.0.1:9100/get_struct_json")
	if err != nil {
		fmt.Println("http.Get err", err)
		return
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("read data err", err)
		return
	}
	fmt.Println(string(data))
}