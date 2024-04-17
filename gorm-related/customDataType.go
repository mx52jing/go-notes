package main

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type CUser struct {
	Gender uint8
	Info   CInfo
}

type CInfo struct {
	Name string
	Age  uint8
}

// Value c就是要向数据库中存储的值
func (c CInfo) Value() (driver.Value, error) {
	byteStr, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	return string(byteStr), nil
}

// Scan value 就是从数据库拿到的值
func (c *CInfo) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("无法解析的数据")
	}
	err := json.Unmarshal(bytes, c)
	if err != nil {
		return err
	}
	return nil
}

func cAutoMigrate() {
	err := DB.AutoMigrate(&CUser{})
	if err != nil {
		return
	}
}

func createCData() {
	u1 := CUser{
		Gender: 0,
		Info: CInfo{
			Name: "张环",
			Age:  19,
		},
	}
	DB.Model(&CUser{}).Create(&u1)
}

func cQueryData() {
	var u CUser
	DB.Model(&CUser{}).First(&u)
	fmt.Printf("u is %v\n", u)
}

func main() {
	//cAutoMigrate()
	//createCData()
	cQueryData()
}
