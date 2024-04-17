package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取当前时间
	// t := time.Now()
	// fmt.Println(t.Format("2006-01-02 15:04:05PM"))
	// fmt.Println("当前时间为", t)
	// // 获取当前时间戳
	// curTimestamp := t.Unix()
	// fmt.Println("当前时间戳为", curTimestamp)

	// // 睡眠两秒
	// time.Sleep(2 * time.Second)
	// d := time.Since(t)
	// fmt.Println(d)
	// t2 := time.Now()
	// // 获取两个时间的间隔
	// timeInterval := t2.Sub(t)
	// fmt.Println("t2和t之前相差的时间按秒来看为:", timeInterval)
	// fmt.Println("t2和t之前相差的时间按毫秒来看为:", timeInterval.Milliseconds())
	// fmt.Println("t2和t之前相差的时间按纳秒来看为:", timeInterval.Microseconds())

	// 计算时间的加法
	// duration := time.Duration(2 * time.Second)
	// t3 := time.Now()
	// fmt.Println("t3时间戳为", t3.Unix())
	// addRes := t3.Add(duration)
	// fmt.Println("t3+duration时间戳为", addRes.Unix())
	//	Year: "2006" "06"
	//	Month: "Jan" "January" "01" "1"
	//	Day of the week: "Mon" "Monday"
	//	Day of the month: "2" "_2" "02"
	//	Day of the year: "__2" "002"
	//	Hour: "15" "3" "03" (PM or AM)
	//	Minute: "4" "04"
	//	Second: "5" "05"
	//	AM/PM mark: "PM"
	t, _ := time.LoadLocation("1998-08-18")
	fmt.Println(t.String())
	// fmt.Println(t.Format("2006/01/02 15:04:05"))
	// fmt.Println(t.Format("2006"))
}