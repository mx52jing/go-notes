package main

import "fmt"


type Sleeper interface {
	sleep()
}
type Dog struct {
	name string
	age  int
}
func (d Dog) sleep() {
	fmt.Printf("%s在睡觉，它今年%d岁了\n", d.name, d.age)
}

type Cat struct {
	name string
	like string
}

func (c Cat) sleep() {
	fmt.Printf("%s 在睡觉，它喜欢%s\n", c.name, c.like)
}

func showInterfaceType(s Sleeper) {
	fmt.Printf("接口s的类型为%T, 值为%v\n", s, s)
}

func emptyInterface(i interface{}) {
	fmt.Printf("i的类型: %T, i的值: %v\n", i, i)
}

func assertType(i interface{}) {
	switch i.(type) {
		case int:
			fmt.Println("int类型")
		case string:
			fmt.Println("string类型")
		case bool:
			fmt.Println("布尔(bool)类型")
		default:
			fmt.Println("未知类型unknown")
	}
}


func main() {
	// dog_one := Dog {
	// 	name: "小黄",
	// 	age: 2,
	// }
	// dog_one.sleep()

	// cat_one := Cat {
	// 	name: "kitty",
	// 	like: "爬",
	// }

	// cat_one.sleep()

	// var dog_two Sleeper
	// dog_two = Dog {
	// 	name: "小灰",
	// 	age: 6,
	// }
	// showInterfaceType(dog_two)
	// dog_two.sleep()
	// name := "张环"
	// emptyInterface(name)
	// age := 22
	// emptyInterface(age)
	// isBoss := true
	// emptyInterface(isBoss)
	// var empty_interface interface {}
	// fmt.Printf("empty_interface 类型为%T, 值为%v\n", empty_interface, empty_interface)
	// var str interface {}
	// str = "张环"
	// fmt.Println(str)
	// str = 88
	// fmt.Println(str)
	// str = true
	// fmt.Println(str)
	// str = []string{"张环", "李朗"}
	// fmt.Println(str)
	// str = map[string]int{
	// 	"name": 2,
	// 	"age": 22,
	// }
	// fmt.Println(str)
	// any_type_value := make([]interface{}, 5)
	// any_type_value[0] = "张环"
	// any_type_value[1] = 22
	// any_type_value[2] = map[int]int{1: 10, 2: 20}
	// any_type_value[3] = []int{2, 22, 222}
	// for idx, value := range any_type_value {
	// 	fmt.Printf("第%d个值为%v\n", idx, value)
	// }
	// assertType(1)
	// assertType("张环")
	// assertType(false)
}