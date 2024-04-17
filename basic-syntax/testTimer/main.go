package main

import (
	"fmt"
	"time"
)

func demo1() {
	timer1 := time.NewTimer(time.Second * 1)

	<-timer1.C
	fmt.Println("timer1 完成")

	timer2 := time.NewTimer(time.Second * 2)

	go func ()  {
		<- timer2.C
	}()
	isTimer2Stop := timer2.Stop()
	fmt.Println("timer stop", isTimer2Stop)
}

func demo2() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)
	go func ()  {
		for {
			select {
				case isDone := <- done:
					fmt.Println("完成了", isDone)
					return
				case t := <- ticker.C:
					fmt.Println("Tick at", t)
			}
		}
	}()
	timer := time.NewTimer(2000 * time.Millisecond)
	<- timer.C
	ticker.Stop()
	done <- true
	fmt.Println("========end==========")
}


func scheduledProcessing() {
	// 声明并初始化一个缓冲区为6的通道，并填充缓冲区，表示有6个任务需要处理
	processing := make(chan int, 6)
	for i := 0; i < 6; i++ {
		processing <- i
	}
	// 关闭通道
	close(processing)
	// 声明一个每200ms执行一次的调度器
	// tick := time.NewTicker(200 * time.Millisecond)
	tick1 := time.Tick(200 * time.Millisecond)

	// 遍历需要处理的任务
	for req := range processing {
		// 每200ms会执行一次释放
		// t := <- tick.C
		t := <- tick1
		fmt.Println("处理了req, req为：", req, t)
	}
}


func burstyLimter() {
	// 初始化一个能容纳3个缓冲的时间通道
	timeLimter := make(chan time.Time, 3)
	// 填充时间通道的缓冲区
	for i := 0; i < 3; i++ {
		timeLimter <- time.Now()
	}

	// 开启一个协程，每200ms向timLimter通道中发送一个time数据
	go func ()  {
		for t:= range time.Tick(200 * time.Millisecond) {
			timeLimter <- t
		}
	}()
	// 声明一个tasks任务通道，并向通道发送数据，填满缓冲区
	tasks := make(chan int, 6)
	for i := 0; i < 6; i++ {
		tasks <- i
	}
	// 关闭tasks通道
	close(tasks)
	for task := range tasks {
		t := <- timeLimter
		fmt.Println("执行task", task, "时间是: ", t)
	}
	// 前3次任务并发执行，后面的每200毫秒处理剩余请求
}



func multiExecCount() {
	limter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		limter <- time.Now()
	}
	

	go func() {
		for ticker := range time.Tick(200 * time.Millisecond) {
			limter <- ticker
		}
	}()

	tasks := make(chan int, 6)
	for i := 0; i < 6; i++ {
		tasks <- i
	}
	close(tasks)

	for task := range tasks {
		t := <- limter
		fmt.Println("执行task", task, "时间是: ", t)
	}
}

func main() {
	// demo1()
	// demo2()
	// scheduledProcessing()
	burstyLimter()
	// multiExecCount()
}