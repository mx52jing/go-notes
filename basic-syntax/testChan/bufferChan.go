package main

// import (
// 	"fmt"
// 	"time"
// )

// // 对于无缓冲channel，发送端和接收端的阻塞问题,发送端在没有准备好之前会阻塞,同样接收端在发送端没有准备好之前会阻塞
// func foo(c chan string) {
// 	<- time.After(time.Second * 2)
// 	fmt.Println("发送端已就绪")
// 	c <- "foo"
// }

// func main() {
// 	chStr := make(chan string)
// 	go func ()  {
// 		foo(chStr)
// 	}()
// 	// 发送端2s后才准备好，所以阻塞在当前位置
// 	fmt.Println("阻塞在当前位置，发送端发送数据后才继续执行")
// 	result := <- chStr
// 	fmt.Printf("end， result的值为%s", result)
// }