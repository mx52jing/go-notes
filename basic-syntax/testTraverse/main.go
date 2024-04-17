package main

import "fmt"

func forAndRangeDifferent() {
	str := "坎坎坷坷9882ajw"
	for i := 0; i < len(str); i++ {
		fmt.Printf("\ti:%d,v:%c\n", i, str[i])
	}
	for i, v := range str {
		fmt.Printf("\ti:%d,v:%c\n", i, v)
	}
}

func main() {
	forAndRangeDifferent()
}