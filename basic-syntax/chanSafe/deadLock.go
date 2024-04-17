package main

import (
	"fmt"
)

func foo(ch chan int) {
	str := <- ch
	fmt.Printf("str is %d", str)
}
func main() {
	ch := make(chan int)
	go func ()  {
		foo(ch)
	}()
	ch <- 98
}