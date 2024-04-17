package main

import "fmt"

func testArr() {
	// var 变量名 [数组长度]数组值类型
	// 数组值为int类型，如果不指定值，默认为0
	// 数组值为string类型，如果不指定值，默认为空字符串
	var arr_init [6]int
	var arr_string [5]string
	arr_no_value := [6]int{}
	fmt.Println(arr_init, arr_string, arr_no_value)

	// 在声明数组时赋值
	// var 变量名 = [数组长度]数组值类型{val1, val2, ...}
	var ages = [2]int{22, 66}
	// 短声明
	short_name := [3]string{"赵", "孙", "李"}
	fmt.Println(ages, short_name)
	// int类型的数组 未赋值的元素默认值为 0
	var some_num = [6]int{22, 66}
	some_str := [6]string{"周", "吴"}
	fmt.Println(some_num, some_str)

	// 为指定的索引赋值
	idx_num := [6]int{1: 22, 5: 88}
	idx_str := [8]string{2: "李", 6: "孙"}
	fmt.Println(idx_num, idx_str)

	// 如果不想写数组长度，可以使用...
	like := [...]string{"篮球","足球","乒乓球", "羽毛球"}
	fmt.Println(like)
}

func testArrLength() {
	names := [2]string{"张环", "李朗"}
	fmt.Println("数组names的长度为:", len(names))
}

func arrayTraverse() {
	likes := [6]string{"篮球", "乒乓球", "羽毛球", "台球", "足球", "网球"}
	for idx, value := range likes {
		fmt.Printf("likes[%d]=%s\n", idx, value)
	}

	for _, value := range likes {
		fmt.Printf("value=%s\n", value)
	}
}


func testArrValue() {
	likes := [6]string{"篮球", "乒乓球", "羽毛球", "台球", "足球", "网球"}
	likes_one := likes;
	likes_one[2] = "溜溜球"
	likes_one[5] = "排球"
	fmt.Println("likes的值为", likes)
	fmt.Println("likes_one的值为", likes_one)
}

func main() {
	// testArr()
	// testArrLength()
	// arrayTraverse()
	testArrValue()
}