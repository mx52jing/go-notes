package main

import "fmt"

func fixedLengthArgs(a, b int, c string) int {
	fmt.Println(c)
	return a + b
}

func moreAargs(age int, isBoss bool, names ...string) {
	fmt.Println("age is", age)
	fmt.Println("isBoss is", isBoss)
	for index, value := range names {
		fmt.Printf("names[%d]=%s\n", index, value)
	}
}

func differentTpyeArgs(arguments ...interface{}) {
	for index, value := range arguments {
		fmt.Printf("第%d个参数的类型为%T\n", index + 1, value)
	}
}

func splitArgs() {
	names := []string{"张环", "李朗"}
	fmt.Println("添加元素之前names为", names)
	names = append(names, []string{"沈韬", "肖豹", "杨方", "仁阔"}...)
	fmt.Println("添加元素之后names为", names)
}

func testA(args ...string) {
	for index, value := range args {
		fmt.Printf("第%d个参数%s\n", index + 1, value)
	}
}

func sum(a, b int) int {
	return a + b
}

func multiReturnVal(a, b int) (int, bool) {
	return a + b, a > b
}

func namesReturnValue(a, b int) (increment int, isLarge bool) {
	increment = a + b
	isLarge = a > b
	return
}

func main() {
	// moreAargs(22, true, "张环", "李朗", "杨方", "仁阔","沈韬", "肖豹")
	// differentTpyeArgs("燕双鹰", 22, 3.1415926, true)
	// splitArgs()
	// testA([]string{"沈韬", "肖豹", "杨方", "仁阔"}...)
	// increment_value := sum(2, 6)
	// fmt.Println(increment_value)
	// increment_value, isLarge := multiReturnVal(2, 6)
	// fmt.Println(increment_value, isLarge)
	increment_value, isLarge := namesReturnValue(2, 6)
	fmt.Println(increment_value, isLarge)
}