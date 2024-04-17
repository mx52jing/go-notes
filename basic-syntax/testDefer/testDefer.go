package main

import (
	"fmt"
)

func deferFunc() {
	fmt.Println("hello")
	defer fmt.Println("defer后跟随的打印")
	fmt.Println("go go go")
}

func deferVaribale() {
	name := "张环"
	defer fmt.Println("defer打印的name", name)

	name = "李朗"
	fmt.Println("name修改后为", name)

	age := []int{22, 66, 88}
	defer fmt.Println("defer打印的age", age)
	age[1] = 99
	fmt.Println("更新后的age为", age)
}

type Reader interface {
	read()
}
type Person struct {
	name	string
	like 	string
}
func (p Person) read() {
	fmt.Printf("%s喜欢%s\n", p.name, p.like)
}

func deferMethod() {
	person := Person {
		name: "张环",
		like: "读书",
	}
	defer person.read()
	fmt.Println("person 值为", person)
}

func deferStack() {
	fmt.Println("start")
	defer fmt.Println("第一个defer打印")
	defer fmt.Println("第二个defer打印")
	defer fmt.Println("第三个defer打印")
	defer fmt.Println("第四个defer打印")
	defer fmt.Println("第五个defer打印")
	defer fmt.Println("第六个defer打印")
	fmt.Println("end")
}

var str string = "张环"

func deferVal() string {
	defer func(){
		str = "李朗"
	}()
	fmt.Println("deferVal 函数内部的 str", str)
	return str
}

func testDeferScope() {
	func (){
		defer fmt.Println("张环")
		defer fmt.Println("李朗")
		defer fmt.Println("沈韬")
	}()
	fmt.Println("-----------------")
	func (){
		defer fmt.Println(22)
		defer fmt.Println(66)
		defer fmt.Println(88)
	}()
	fmt.Println("=================")
}

func testDeferReturn() (num int) {
	num = 10
	defer func(){
		num += 5
	}()
	return 2
}

func testDeferPanic() {
	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// defer fmt.Println(3)
	// panic("------------")
	// defer fmt.Println(5)
	defer func (){
		fmt.Println("1223456")
	}()
	panic("9999999")
}

func testCatchPanic() {
	defer func (){
		if err := recover(); err != nil {
			fmt.Println( "程序崩溃原因：", err)
		}
	}()
	panic("来，崩溃吧")
}

func main() {
	// deferFunc()
	// deferVaribale()
	// deferMethod()
	// deferStack()
	// str_val := deferVal();
	// fmt.Println("str_val", str_val)
	// fmt.Println("deferVal函数外部的str", str)
	// testDeferScope()
	// num := testDeferReturn()
	// fmt.Println(num)
	// testDeferPanic()
	testCatchPanic()
}