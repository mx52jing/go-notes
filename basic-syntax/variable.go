package main

import "fmt"

func variable() {
	// 声明变量的方式
	var a int = 1
	var b, c int = 2, 3 // 声明两个变量
	var (               // 声明多个变量
		d    int = 5
		e, f int = 6, 7
	)
	var g = 8
	fmt.Println("【使用var声明的变量】", a, b, c, d, e, f, g)
	// 另外一种声明变量的方式
	h := 9
	i := "I am I"
	fmt.Println("【使用:= 声明的变量】", h, i)
	// 常量
	const (
		A0 = 0
		A1 = 22
		A2 = 88
	)
	fmt.Println("【A0, A1, A2】:", A0, A1, A2)
	// iota 默认为0，每行默认+1 只能用在const中
	const (
		B0 = iota
		B1
		B2
	)
	fmt.Println("【B0, B1, B2】:", B0, B1, B2)
	const (
		C0 = iota + 1
		C1
		C2
	)
	fmt.Println("【C0, C1, C2】:", C0, C1, C2)
	// 当不想使用某个值时，可以使用_ 代替
	const (
		D0 = iota * 2
		_
		D2
		_
		D5
	)
	fmt.Println("【D0, D2, D5】:", D0, D2, D5)
}
