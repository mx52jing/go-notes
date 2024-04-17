package main

import (
	"fmt"
	"time"
)

func printText() {
	fmt.Println("hello world")
}


func loopPrint(num int) {
	for i := 0; i < 2; i++ {
		fmt.Println(num)
		time.Sleep(time.Millisecond)
	}
}

func main() {
	// 开启一个协程执行 printText 函数
	// go printText()
	// 使主协程休眠 1 秒
	// time.Sleep(time.Second * 1)
	// fmt.Println("printText函数之后")
	go loopPrint(2)
	go loopPrint(6)
	// time.Sleep(time.Second)
	fmt.Println("loopPrint函数之后")
}