package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 常规写入
func demo1() {
	// 如果文件不存在就创建并写入内容
	file, err := os.OpenFile("2.txt", os.O_CREATE|os.O_RDWR, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println("os.OpenFile err", err)
		return
	}
	// func (f *File) WriteString(s string) (n int, err error)
	n, err := file.WriteString("树倒猢狲散\n")
	// func (f *File) Write(b []byte) (n int, err error) 
	n, err = file.Write([]byte("啊是, 卡手机报价"))
	if err != nil {
		fmt.Println("file write err", err)
		return
	}
	fmt.Println(n)
}

// 快速写
func demo2() {
	// func WriteFile(name string, data []byte, perm FileMode)
	err := os.WriteFile("3.txt", []byte("齐天大圣孙悟空，hello"), 0666)
	if err != nil {
		fmt.Println("os.WriteFile", err)
		return
	}	
}

// 带缓冲写，一般用于大文件
func demo3() {
	file, err := os.OpenFile("5.txt", os.O_CREATE|os.O_RDWR, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println("os.OpenFile err", err)
		return
	}
	buf := bufio.NewWriter(file)
	for i := 0; i <= 5; i++ {
		buf.WriteString("go语言学习go go go\n")
	}
	buf.Flush()
}

func binaryFileOperation() {
	// 打开1.jpg文件
	imgFile, imgErr := os.OpenFile("1.jpg", os.O_RDONLY, 0777)
	defer imgFile.Close() // 读取结束后，需要关闭文件
	if imgErr != nil {
		fmt.Println("os.OpenFile", imgErr)
		return
	}
	copyImgFile, copyImgErr := os.OpenFile("fengjing.jpg", os.O_CREATE|os.O_RDWR, 0666)
	if copyImgErr != nil {
		fmt.Println("copyImgErr", imgErr)
		return
	}
	reader := bufio.NewReader(imgFile)
	writer := bufio.NewWriter(copyImgFile)
	defer copyImgFile.Close()
	for {
		byte, err := reader.ReadByte() // 按字节读取
		if err == io.EOF {
			fmt.Println("======读取完毕======")
			break
		}
		writer.WriteByte(byte)
	}
	writer.Flush()
}


func binaryFileOperation2() {
	readFile, err := os.ReadFile("1.jpg")
	if err != nil {
		fmt.Println("open readFile err", err)
		return
	}
	writeErr := os.WriteFile("write.jpg", readFile, 0666)
	if writeErr != nil {
		fmt.Println("open writeFile err", writeErr)
		return
	}
	fmt.Println("文件读取完毕")
}

func copyFile() {
	readFile, readFileErr := os.Open("1.jpg")
	defer readFile.Close()
	if readFileErr != nil {
		fmt.Println("Open file err", readFileErr)
		return
	}
	writeFile, writeFileErr := os.Create("copy.jpg")
	if writeFileErr != nil {
		fmt.Println("writeFile err", writeFileErr)
		return
	}
	n, copyErr := io.Copy(writeFile, readFile)
	if copyErr != nil {
		fmt.Println("copy err", copyErr)
		return
	}
	fmt.Println("n", n, "copy success")
}


func dirOperation() {
	dirEntry, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("dirEntry err", err)
		return
	}
	for _, entry := range dirEntry {
		fmt.Println(entry.IsDir(), entry.Name())
		fileInfo, _ := entry.Info()
		fmt.Println(fileInfo.Size())
	}
}

func main() {
	// demo1()
	// demo2()
	// demo3()
	// binaryFileOperation()
	// binaryFileOperation2()
	// copyFile()
	dirOperation()
}