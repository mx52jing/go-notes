package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	StudentId int
	Age int
	UserName string
	IsMonitor bool
}

type ClassOne struct {
	Id int
	Students []Student
}

func main() {
	xiaoMing := Student{ 1, 20, "小明", false }
	xiaoHong := Student{ 2, 22, "小红", true }
	classOneInstance := ClassOne{
		Id: 1,
		Students: []Student { xiaoMing, xiaoHong },
	}
	// func json.Marshal(v any) ([]byte, error)
	// 将结构体转换为二进制流
	bytes, marshalErr := json.Marshal(classOneInstance)
	if marshalErr != nil {
		fmt.Println("json 序列化失败", marshalErr)
		return
	}
	fmt.Println(bytes)
	fmt.Println(string(bytes))

	var cls ClassOne
	// func json.Unmarshal(data []byte, v any) error
	unmarshalErr := json.Unmarshal(bytes, &cls)
	if unmarshalErr != nil {
		fmt.Println("json 反序列化失败", unmarshalErr)
		return
	}
	fmt.Printf("cls的值为 %v\n", cls)
}