package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	const N = 10
// 	mMap := make(map[int]int)
// 	wg := &sync.WaitGroup{}
// 	mu := &sync.Mutex{}
// 	wg.Add(N)
// 	for i := 0; i < N; i++ {
// 		go func(i int){
// 			defer wg.Done()
// 			mu.Lock()
// 			fmt.Println(i, "<<<<<<i>>>>>>")
// 			mMap[i] = i
// 			mu.Unlock()
// 		}(i)
// 	}
// 	wg.Wait()
// 	fmt.Println(len(mMap), mMap)
// }