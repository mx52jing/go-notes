package main

import (
	"fmt"
)

func createPointer() {
	name := "张环"
	name1 := &name

	fmt.Println(name1) // 0x14000010270

	str := new(string)
	*str = "蛇灵"

	str1 := *str

	fmt.Println("str:", str, *str)
	fmt.Println("str1:", str1)

	// 定义普通变量
	title := "蛇灵-闪灵"
	// 定义指针变量
	var title1 *string
	// 将指针的值指向普通变量的内存地址
	title1 = &title
	fmt.Println(*title1, title)


	x := "蛇灵-血灵"
	y := &x

	fmt.Println("x =>", x)
	fmt.Println("*y => ", *y)
	fmt.Println("&x =>", &x)
	fmt.Println("y =>", y)
}

func pointerType() {
	name := "魔灵"
	age := 22
	isBoss := false
	level := 32.86
	fmt.Printf("type of &name is: %T \n", &name)
	fmt.Printf("type of &age is: %T \n", &age)
	fmt.Printf("type of &isBoss is: %T \n", &isBoss)
	fmt.Printf("type of &level is: %T \n", &level)
}

func zreoPointer() {
	name := "元芳"
	var mentor *string
	fmt.Println("初始化未赋值时为：", mentor)
	mentor = &name
	fmt.Println("赋值之后为:", mentor)
}

func changePointer(age *int) {
	fmt.Println("传递进来的age为：", age)
	*age = 88
}

func testChangePointer() {
	age_one := 22
	age_two := &age_one
	fmt.Println("执行changePointer函数之前age_two的值为：", *age_two)
	changePointer(age_two)
	fmt.Println("执行changePointer函数之后age_two的值为：", *age_two)
	// 由于修改了指指针指向的变量的值，所以原来的变量age_one也被修改了
	fmt.Println("执行changePointer函数之后age_one的值为：", age_one)
}


func changeArrayBySlice(value []string) {
	value[1] = "肖豹"
}

func changeArrayByPointer(value *[3]int) {
	(*value)[2] = 99
}

func testChangeArray() {
	names := [3]string{"张环", "李朗", "沈韬"}
	fmt.Println("执行changeArrayBySlice函数前names的值为:", names)
	changeArrayBySlice(names[:])
	fmt.Println("执行changeArrayBySlice函数后names的值为:", names)

	ages := [3]int{22, 66, 88}
	// 执行changeArrayByPointer函数前ages的值为: [22 66 88]
	fmt.Println("执行changeArrayByPointer函数前ages的值为:", ages)
	changeArrayByPointer(&ages)
	// 执行changeArrayByPointer函数前ages的值为: [22 66 99]
	fmt.Println("执行changeArrayByPointer函数前ages的值为:", ages)
}

func pointerCalc() {
	// name := "张环"
	// colleague := &name
	// fmt.Println(colleague++)
}

func main() {
	// createPointer()
	// pointerType()
	// zreoPointer()
	// testChangePointer()
	// testChangeArray()
	pointerCalc()
}