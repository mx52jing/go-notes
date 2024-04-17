package main

import (
	"fmt"
	"sync"
)

type Balance struct {
	count int
	mutex sync.Mutex
	rwMutex sync.RWMutex
}

func (b *Balance) deposite(num int) {
	b.mutex.Lock()
	defer b.mutex.Unlock();
	b.count += num
}


func (b *Balance) deposite2(num int) {
	b.rwMutex.Lock()
	defer b.rwMutex.Unlock()
	b.count += num
}

func (b *Balance) getBalance() int {
	b.rwMutex.RLock()
	defer b.rwMutex.RUnlock();
	return b.count
}

func syncTest() {
	var wg sync.WaitGroup
	num := 10000
	balance := &Balance{}
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func ()  {
			balance.deposite(num)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Balance count", balance.count)
}

func testMutex() {
	var wg sync.WaitGroup
	num := 10000
	balance := &Balance{}
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func ()  {
			balance.deposite(num)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Balance count", balance.getBalance())
}


func testRWMutex() {
	var wg sync.WaitGroup
	num := 10000
	balance := &Balance{}
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func ()  {
			balance.deposite2(num)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Balance count", balance.getBalance())
}

func main(){
	// syncTest()
	// testMutex()
	testRWMutex()
}