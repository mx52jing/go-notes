package main

import "fmt"

func createSlice() {
	var slice_one []int
	fmt.Println(slice_one)

	slice_two := []int{}
	fmt.Println(slice_two)

	var slice_three = []string{"张环", "李朗"}
	fmt.Println(slice_three)

	slice_four := []int{22, 66, 88}
	fmt.Println(slice_four)


	slice_five := []int{1: 22, 5: 88}
	fmt.Println(slice_five)

	slice_six := make([]int, 3, 3)
	fmt.Println(slice_six)

	slice_seven := make([]int, 2, 6)
	fmt.Println(slice_seven)
	slice_seven = append(slice_seven, 2)
	slice_seven = append(slice_seven, 4)
	slice_seven = append(slice_seven, 6)
	slice_seven = append(slice_seven, 8)
	fmt.Println(slice_seven, len(slice_seven), cap(slice_seven), slice_seven[:])

	slice_eight := slice_seven[2:]
	fmt.Println(slice_eight)

	slice_eight[3] = 88

	fmt.Println(slice_seven, slice_eight)
}

func testLenCap() {
	age := make([]int, 2, 6)
	fmt.Println(age, len(age), cap(age))
	
	fmt.Println(age[7])
	age[8] = 0
}

func slice_no_value() {
	var like []string
	fmt.Println(like, like == nil)
}

func testGetValue() {
	names := []string{"张环", "李朗", "沈韬", "肖豹", "杨方", "仁阔"}

	fmt.Println(names[:])
	fmt.Println(names[2:])
	fmt.Println(names[2:5])
	fmt.Println(names[:5])
}

func sliceAppend() {
	age := []int{18}
	fmt.Println(age, len(age), cap(age))

	age = append(age, 22)
	fmt.Println(age, len(age), cap(age))

	age = append(age, 33, 66)
	fmt.Println(age, len(age), cap(age))

	age = append(age, []int{2, 5, 8, 0}...)
	fmt.Println(age, len(age), cap(age))
}


func multidimensionalSlice(){
	str := [][]string{ {"包子", "饺子"},{"硬币", "鞭炮"}}
	fmt.Println(str)
}

var ages []int
func argumentSlice(a []int) {
	fmt.Println("ages is", a);
	a[1] = 188
	a = append(a, 99)
	fmt.Println("ages is", a);
}

func main() {
	// createSlice()
	// testLenCap()
	// slice_no_value()
	// testGetValue()
	// sliceAppend()
	ages = []int{22, 66, 88}
	// multidimensionalSlice()
	argumentSlice(ages)
	fmt.Println("==========ages", ages)
}