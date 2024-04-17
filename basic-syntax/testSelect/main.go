package main

import (
	"fmt"
	"time"
)

func demo1() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	ch3 := make(chan string, 1)
	ch1 <- "张环"
	ch2 <- "李朗"
	ch3 <- "沈韬"
	select {
	case msg1 := <- ch1:
		fmt.Println("从ch1收到数据", msg1)
	case msg2 := <- ch2:
		fmt.Println("从ch2收到数据", msg2)
	case msg3 := <- ch3:
		fmt.Println("从ch3收到数据", msg3)
	default:
		fmt.Println("not receive data")
	}
}

func task1(ch chan string) {
	time.Sleep(time.Microsecond * 2)
	ch <- "张环"
}	

func task2(ch chan string) {
	time.Sleep(time.Microsecond * 3)
	ch <- "李朗"
}	
func task3(ch chan string) {
	time.Sleep(time.Microsecond * 1)
	ch <- "沈韬"
}	
func demo2() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	ch3 := make(chan string, 1)
	go task1(ch1)
	go task2(ch2)
	go task3(ch3)
	select {
		case str1 := <- ch1:
			fmt.Println("ch1 received", str1)
		case str2 := <- ch2:
			fmt.Println("ch2 received", str2)
		case str3 := <- ch3:
			fmt.Println("ch3 received", str3)
	}
}

func demo3() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	ch3 := make(chan string, 1)
	// ch1 <- "张环"
	// ch2 <- "李朗"
	// ch3 <- "沈韬"
	select {
	case msg1 := <- ch1:
		fmt.Println("从ch1收到数据", msg1)
	case msg2 := <- ch2:
		fmt.Println("从ch2收到数据", msg2)
	case msg3 := <- ch3:
		fmt.Println("从ch3收到数据", msg3)
	// default:
	// 	fmt.Println("not receive data")
	}
}

func emptySelect() {
	select {

	}
}

func makeTimeout(ch chan int) {
	time.Sleep(time.Microsecond)
	ch <- 1
}

func testTimeout() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	ch3 := make(chan string, 1)
	timeoutChan := make(chan int)
	go makeTimeout(timeoutChan)
	select {
	case msg1 := <- ch1:
		fmt.Println("从ch1收到数据", msg1)
	case msg2 := <- ch2:
		fmt.Println("从ch2收到数据", msg2)
	case msg3 := <- ch3:
		fmt.Println("从ch3收到数据", msg3)
	case <- timeoutChan:
		fmt.Println("超时")
	}
}

func delay(ch chan<- string) {
	time.Sleep(time.Second * 2)
	ch <- "张环"
}

func testTimeout1() {
	ch := make(chan string)
	go delay(ch)
	select {
		case msg := <-ch:
			fmt.Println("msg接收到的数据", msg)
		case <-time.After(1 * time.Second):
			fmt.Println("超时处理了")
	}
}

func main() {
	// demo1()
	// demo2()
	// demo3()
	// emptySelect()
	// fatal error: all goroutines are asleep - deadlock!
	// testTimeout()
	testTimeout1()
}