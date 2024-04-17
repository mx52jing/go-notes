package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func createChan() {
// 	// 声明一个string类型的通道，并未初始化
// 	var name chan string
// 	fmt.Println("声明但并未初始化的chan的值为", name)
// 	// 初始化chan name
// 	name = make(chan string)
// 	fmt.Println("初始化后的chan的值为", name)

// 	// 声明并初始化一个chan
// 	age := make(chan int)
// 	fmt.Println("声明并初始化一个chan，值为", age)
// }

// func setChanValue(val chan string) {
// 	defer close(val)
// 	val <- "张环"
// 	val <- "李朗"
// }

// func sendAndReceiveData() {
// 	fmt.Println("start")
// 	str := make(chan string)
// 	go setChanValue(str)
// 	for {
// 		val, isClose := <- str
// 		if(!isClose) {
// 			fmt.Println("通道关闭")
// 			break
// 		}
// 		fmt.Println("从通道中获取的数据", val)
// 	}
// 	// receiveValue, isCose := <- str
// 	// fmt.Println("接收到通道的数据receiveValue为", receiveValue, isCose)
// 	// fmt.Println("end")
// 	// _, ok := <- str
// 	// // 检测通道是否已经关闭，false 已关闭 true 未关闭
// 	// fmt.Println(ok)
// }

// func noBuffersizeChannel() {
// 	ch := make(chan string, 3)
// 	ch <- "张环"
// 	ch <- "李朗"
// 	fmt.Println("==============")
// }

// func chanLenCap() {
// 	str := make(chan string, 3)
// 	fmt.Println("初始化后：")
// 	fmt.Println("str的len", len(str))
// 	fmt.Println("str的cap", cap(str))
// 	str <- "张环"
// 	fmt.Println("放入一个数据后：")
// 	fmt.Println("str的len", len(str))
// 	fmt.Println("str的cap", cap(str))
// 	str <- "李朗"
// 	fmt.Println("放两个数据后：")
// 	fmt.Println("str的len", len(str))
// 	fmt.Println("str的cap", cap(str))
// 	<- str
// 	fmt.Println("取出一个数据后：")
// 	fmt.Println("str的len", len(str))
// 	fmt.Println("str的cap", cap(str))
// }

// func oneWayChannel() {
// 	// 定义可读通道
// 	// type ReaderChan <-chan string
// 	type ReaderChan = <-chan string
// 	// 定义可写通道
// 	type WriterChan = chan<- string

// 	// 声明并初始化一个双向通道
// 	bothChan := make(chan string)

// 	go func(){
// 		// 只写通道
// 		// 这种定义方法也可以
// 		// var writerChan WriterChan = bothChan
// 		writerChan := WriterChan(bothChan)
// 		fmt.Println("准备写入数据")
// 		writerChan <- "张环、李朗"
// 	}()
// 	go func(){
// 		// 只读通道
// 		// 这种定义方法也可以
// 		// var readerChan ReaderChan = bothChan
// 		readerChan := ReaderChan(bothChan)
// 		fmt.Println("准备读取数据")
// 		msg := <- readerChan
// 		fmt.Println("读取到的数据为", msg)
// 	}()
// 	time.Sleep(time.Second * 1)
// 	fmt.Println("===================")
// }

// func traverseChannel() {
// 	ch := make(chan int)
// 	// 生产者，向通道中发送数据
// 	go func() {
// 		for i := 1; i < 7; i++ {
// 			ch <- i
// 		}
// 		close(ch) // 关闭通道
// 	}()
// 	// 消费者，从通道中接收数据
// 	for num := range ch {
// 		fmt.Println("接收到的数据为=>", num)
// 	}
// }

// func increment(num chan string, idx *int) {
// 	num <- "蛇灵"
// 	*idx += 1
// 	<- num
// }

// // 使用容量为 1 的通道可以达到锁的效果
// func useChannelLock() {
// 	num := make(chan string, 1)
// 	idx := 0
// 	fmt.Println("------start------")
// 	for i := 0; i < 100000; i++ {
// 		go increment(num, &idx)
// 	}
// 	time.Sleep(time.Second)
// 	fmt.Println("idx值为", idx)
// 	fmt.Println("------end------")
// }

// func task(id int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Printf("执行第%d条任务\n", id)
// }

// func basicChan() {
// 	msg := make(chan string)
// 	go func ()  {
// 		msg <- "张环"
// 	}()
// 	str := <- msg
// 	fmt.Println("str", str)
// }

// func hasBufferSizeChan() {
// 	str := make(chan string, 2)
// 	str <- "z"
// 	str <- "h"
// 	str <- "a"
// 	fmt.Println(<- str)
// 	fmt.Println(<- str)
// 	fmt.Println(<- str)
// }

// func twoChan() {
// 	c1 := make(chan string)
// 	c2 := make(chan string)
// 	go func ()  {
// 		time.Sleep(time.Millisecond * 2)
// 		c1 <- "hello"
// 	}()
// 	go func ()  {
// 		time.Sleep(time.Millisecond * 10)
// 		c2 <- "go"
// 	}()
// 	for i := 0; i < 2; i++ {
// 		select {
// 			case reveiveC1 := <- c1:
// 				fmt.Println("c1", reveiveC1)
// 			case reveiveC2 := <- c2:
// 				fmt.Println("c2", reveiveC2)
// 		}
// 	}
// }

// func worker(id int, job <-chan int, res chan<- int) {
// 	for j := range job {
// 		fmt.Println("worker", id, "start job", j)
// 		time.Sleep(time.Second)
// 		fmt.Println("worder", id, "finish job", j)
// 		res <- j * 2
// 	}
// }

// func testPool() {
// 	total := 5;
// 	job := make(chan int, total)
// 	res := make(chan int, total)
// 	for i := 1; i <= 3; i++ {
// 		go worker(i, job, res)
// 	}
// 	for i := 1; i <= total; i++ {
// 		job <- i
// 	}
// 	close(job)
// 	for a := 1; a <= total; a++ {
// 		<-res
// 	}
// }

// func main() {
// 	// createChan()
// 	// sendAndReceiveData()
// 	// noBuffersizeChannel()
// 	// chanLenCap()
// 	// oneWayChannel()
// 	// traverseChannel()
// 	// useChannelLock()
// 	// var wg sync.WaitGroup
// 	// wg.Add(6)
// 	// // 启动六个子 goroutine
// 	// go task(1, &wg)
// 	// go task(2, &wg)
// 	// go task(3, &wg)
// 	// go task(4, &wg)
// 	// go task(5, &wg)
// 	// go task(6, &wg)
// 	// // 等待所有子 goroutine 完成
// 	// wg.Wait()
// 	// // 所有子 goroutine 完成后，执行后续操作
// 	// fmt.Println("所有子goroutine已完成任务，可以继续执行后续操作")
// 	// basicChan()
// 	// hasBufferSizeChan()
// 	// twoChan()
// 	testPool()
// }