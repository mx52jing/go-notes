package main

// import (
// 	"fmt"
// 	"sync"
// )

// type WalletV1 struct {
// 	balance int
// }

// func (w *WalletV1) deposite(num int) {
// 	w.balance += num
// }

// func (w *WalletV1) getBalance() int {
// 	return w.balance
// }

// func saveMoneyV1() {
// 	num := 10000;
// 	walletV1 := &WalletV1{}
// 	var wg sync.WaitGroup
// 	wg.Add(num)
// 	for i := 0; i < num; i++ {
// 		go func ()  {
// 				walletV1.deposite(num)
// 				wg.Done()
// 		}()
// 	}
// 	wg.Wait()
// 	fmt.Println("余额:", walletV1.getBalance())
// }

// type WalletV2 struct {
// 	balance int
// 	m				sync.Mutex
// }

// func (w *WalletV2) deposite(num int) {
// 	w.m.Lock()
// 	defer w.m.Unlock()
// 	w.balance += num
// }

// func (w *WalletV2) getBalance() int {
// 	return w.balance
// }

// func saveMoneyV2() {
// 	num := 10000;
// 	walletV2 := &WalletV2{}
// 	var wg sync.WaitGroup
// 	wg.Add(num)
// 	for i := 0; i < num; i++ {
// 		go func ()  {
// 				walletV2.deposite(num)
// 				wg.Done()
// 		}()
// 	}
// 	wg.Wait()
// 	fmt.Println("余额:", walletV2.getBalance())
// }

// type WalletV3 struct {
// 	balance int
// 	m				sync.Mutex
// 	rw      sync.RWMutex
// }

// func (w *WalletV3) deposite(num int) {
// 	w.rw.Lock()
// 	defer w.rw.Unlock()
// 	w.balance += num
// }

// func (w *WalletV3) getBalance() int {
// 	w.rw.RLock()
// 	defer w.rw.RUnlock()
// 	return w.balance
// }

// func saveMoneyV3() {
// 	num := 10000;
// 	walletV3 := &WalletV3{}
// 	var wg sync.WaitGroup
// 	wg.Add(num)
// 	for i := 0; i < num; i++ {
// 		go func ()  {
// 				walletV3.deposite(num)
// 				wg.Done()
// 		}()
// 	}
// 	wg.Wait()
// 	fmt.Println("余额:", walletV3.getBalance())
// }
// func main() {
// 	// saveMoneyV1()
// 	// saveMoneyV2()
// 	saveMoneyV3()
// }