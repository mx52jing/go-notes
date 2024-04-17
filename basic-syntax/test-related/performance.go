package test_related

import (
	"encoding/json"
	"github.com/bytedance/sonic"
)

type Student struct {
	Name   string
	Age    uint8
	Gender string
}

type Class struct {
	Id       string
	Students []Student
}

var (
	s1 = Student{Name: "闪灵", Age: 26, Gender: "male"}
	s2 = Student{Name: "魔灵", Age: 28, Gender: "male"}
	s3 = Student{Name: "血灵", Age: 18, Gender: "female"}
	C1 = Class{
		Id:       "class_c1",
		Students: []Student{s1, s2, s3},
	}
)

func handleJsonByJsonLibrary() (Class, error) {
	marshal, _ := json.Marshal(C1)
	//if err != nil {
	//	return Class{}, err
	//}
	var jsonC2 Class
	json.Unmarshal(marshal, &jsonC2)
	//if err != nil {
	//	return Class{}, err
	//}
	return jsonC2, nil
}

func handleJsonBySonicLibrary() (Class, error) {
	marshal, _ := sonic.Marshal(C1)
	//if err != nil {
	//	return Class{}, err
	//}
	var sonicC2 Class
	sonic.Unmarshal(marshal, &sonicC2)
	//if err != nil {
	//	return Class{}, err
	//}
	return sonicC2, nil
}
