package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"path/filepath"
// 	"strconv"
// )

// func main() {
// 	// 说明: 思路是遍历根目录下所有的文件，如果是根目录下的文件夹就开一个线程，如果是文件就直接+1，
// 	// 根目录下的文件夹里的文件夹不会单独再开线程，会使用原线程一直查询下去，直到查完所有文件
// 	searchRoot2("./")
// }

// var filesCount2 = 0

// func search2(path string, c chan int) {
// 	files, err := ioutil.ReadDir(path)
// 	if err == nil {
// 		for _, file1 := range files {
// 			name := file1.Name()
// 			if file1.IsDir() {
// 				fmt.Println("进入文件夹: " + path + string(filepath.Separator) + name)
// 				search2(path+string(filepath.Separator)+name, c)
// 			} else {
// 				fmt.Println("找到文件: " + path + string(filepath.Separator) + name)
// 				c <- 1
// 			}
// 		}
// 	} else {
// 		fmt.Println("error: ", err)
// 	}
// }

// var chans2 []chan int

// // 把查询根目录下的文件，每一个开一个协程去查询
// func searchRoot2(path string) {
// 	dir, err := ioutil.ReadDir(path)
// 	if err == nil {
// 		for _, info := range dir {
// 			if info.IsDir() {
// 				ci := make(chan int)
// 				chans2 = append(chans2, ci)
// 				go search2(path+string(filepath.Separator)+info.Name(), ci)
// 			} else {
// 				filesCount2++
// 			}
// 		}
// 		fmt.Println("goroutine数量: " + strconv.Itoa(len(chans2)))
// 		for {
// 			i := 0
// 			for _, item := range chans2 {
// 				select {
// 				case data, open := <-item:
// 					if open {
// 						filesCount2 += data
// 					} else {
// 						fmt.Printf("通道关闭+1")
// 						//sw.Done()
// 						i++
// 					}
// 					//default:
// 					//fmt.Printf("继续读取")
// 				}
// 			}
// 			if i == len(chans2) {
// 				fmt.Printf("所有channel都被关闭,执行结束: %d", filesCount2)
// 				break
// 			}
// 		}
// 	}
// }