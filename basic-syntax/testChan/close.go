package main

import (
	"fmt"
	"time"
)

func main() {
	chan1 := make(chan int, 0)
	go func() {
		time.Sleep(5 * time.Second)
		<-chan1
	}()
	chan1 <- 1
	fmt.Println("cover")
	//c := make(chan int, 10)
	//
	//for i := 0; i < 10; i++ {
	//	c <- i
	//}
	//close(c)
	//// for val := range c {
	//// 	fmt.Println(val)
	//// }
	//for {
	//	val, ok := <-c
	//	if !ok {
	//		fmt.Println("没有值可取", val, ok)
	//		break
	//	}
	//	fmt.Println(val, ok)
	//}
	//fmt.Println("end")
}
