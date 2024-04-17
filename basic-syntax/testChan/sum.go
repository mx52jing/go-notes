package main

// import "fmt"

// func sum(arr []int, c chan int) {
// 	res := 0
// 	for _, val := range arr {
// 		res += val
// 	}
// 	c <- res
// }

// func main() {
// 	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	count := make(chan int)
// 	go sum(arr[0:len(arr)/2], count)
// 	go sum(arr[len(arr)/2:], count)
// 	x, y := <- count, <- count
// 	fmt.Println(x, y, x + y)
// }