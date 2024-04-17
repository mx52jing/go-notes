package main

import "fmt"

func baseGoTo() {
	fmt.Println("start goto")
	goto skip
	fmt.Println("中间打印")
	skip: fmt.Println("直接调转到这里")
}

func testDefineVariable() {
	fmt.Println("start goto")
	// goto define
	fmt.Println("中间打印")
	name := "sdf"
	fmt.Println(name)
	// define: fmt.Println("直接调转到这里")	
}

func main() {
	// baseGoTo()
	testDefineVariable()
}