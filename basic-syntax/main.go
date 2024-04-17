package main // 声明main 包

import (
	"fmt"
	"time"
)

func testFor() {
	i := 1
	for i <= 3 {
		fmt.Println("i的值为：", i)
		i++
	}

	for f := 1; f <= 3; f++ {
		fmt.Println("f的值为：", f)
	}

	k := 1
	for {
		if k > 3 {
			break
		}
		fmt.Println("k的值为：", k);
		k++
	}
}

func testIfElse() {
	i := 6
	if i > 5 {
		fmt.Println("i的值大于5，i为：", i)
	}

	if f := 12; f >= 18 {
		fmt.Println("大于等于18岁，成年了")
	} else if f > 10 {
		fmt.Println("大于10岁，小于18岁")
	} else {
		fmt.Println("不到10岁")
	}
	// fmt.Println(f) 在if后面声明的变量，外部是拿不到的，只有if else 语句能拿到
}

func testSwitch() {
	switch time.Now().Weekday() {
		case time.Monday:
			fmt.Println("周一")
		case time.Tuesday:
			fmt.Println("周二")
			// fallthrough
		case time.Saturday, time.Sunday:
			fmt.Println("周末")
		default:
			fmt.Println("周三到周五")
	}
}

func main() {
	// testFor()
	// testIfElse()
	// testSwitch()
	const s string = "hello go"
	var (
		name, age = "张三", 22
	)
	fmt.Println(s, name, age)
}