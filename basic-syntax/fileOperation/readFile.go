package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"os"
// )

// // 文件读取
// // Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
// // O_RDONLY int = syscall.O_RDONLY // open the file read-only.
// // O_WRONLY int = syscall.O_WRONLY // open the file write-only.
// // O_RDWR   int = syscall.O_RDWR   // open the file read-write.
// // // The remaining values may be or'ed in to control behavior.
// // O_APPEND int = syscall.O_APPEND // append data to the file when writing.
// // O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
// // O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
// // O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
// // O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.

// // O_RDONLY int = syscall.O_RDONLY // 只读
// // O_WRONLY int = syscall.O_WRONLY // 只写
// // O_RDWR   int = syscall.O_RDWR   // 读写
// // O_APPEND int = syscall.O_APPEND // 追加
// // O_CREATE int = syscall.O_CREAT  // 如果不存在就创建
// // O_EXCL   int = syscall.O_EXCL   // 文件必须不存在
// // O_SYNC   int = syscall.O_SYNC   // 同步io
// // O_TRUNC  int = syscall.O_TRUNC  // 打开时清空文件

// // 一次性全部读取
// // func ReadFile(name string) ([]byte, error)
// func demo1() {
// 	file, err := os.ReadFile("1.txt");
// 	if err != nil {
// 		fmt.Println("os.ReadFile err", err)
// 		return
// 	}
// 	fmt.Println(string(file))
// }

// // 读取指定字节长度
// // func OpenFile(name string, flag int, perm FileMode) (*File, error)
// // 中文汉字和中文符号为3个字节，英文字母和英文符号为1个字节，
// func demo2() {
// 	file, err := os.OpenFile("1.txt", os.O_RDONLY, 0777)
// 	if err != nil {
// 		fmt.Println("os.OpenFile", err)
// 		return
// 	}
// 	str := make([]byte, 9)
// 	n, err := file.Read(str)
// 	if err != nil {
// 		fmt.Println("file.Read", err)
// 		return
// 	}
// 	fmt.Println("n", n)
// 	fmt.Println("读到的str为:", string(str))
// }

// // 读取片段
// // func (f *File) Seek(offset int64, whence int) (ret int64, err error)
// // offset 表示指针的偏移量
// // whence 表示指针移动的方式
// // 0 从数据头部开始移动指针；1 从数据的当前位置开始移动指针；2 从数据的尾部开始移动指针
// func demo3() {
// 	file, err := os.OpenFile("1.txt", os.O_RDONLY, 0777)
// 	if err != nil {
// 		fmt.Println("os.OpenFile", err)
// 		return
// 	}
// 	str := make([]byte, 9)
// 	file.Seek(9, 0)
// 	n, err := file.Read(str)
// 	if err != nil {
// 		fmt.Println("file.Read", err)
// 		return
// 	}
// 	fmt.Println("n", n)
// 	fmt.Println("读到的str为:", string(str))
// }

// // 带缓存读取
// func demo5() {
// 	file, err := os.OpenFile("1.txt", os.O_RDONLY, 0777)
// 	defer file.Close() // 读取结束后，需要关闭文件
// 	if err != nil {
// 		fmt.Println("os.OpenFile", err)
// 		return
// 	}
// 	reader := bufio.NewReader(file)
// 	for {
// 		line, _, err := reader.ReadLine() // 按行读，每次读取一行
// 		if err == io.EOF {
// 			fmt.Println("======读取完毕======")
// 			break
// 		}
// 		fmt.Println(string(line))
// 	}
// }

// func main() {
// 	// demo1()
// 	// demo2()
// 	// demo3()
// 	// demo5()
// }