package main

import (
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

type User struct {
	UserName string `json:"user_name"`
	Password string `json:"password,omitempty"`
	Age      int    `json:"-"`
	Gender   int
	email    string
}

type OtherTagOptions string // json tag中的其他选项

/*
*
解析tag 解析出name 和 tag options
例如 `json:user_name,omitempty` 解析出name就是 user_name, tag options 就是omitempty
*/
func parseTag(tag string) (string, OtherTagOptions) {
	// 判断tag中是否有逗号 ,
	commaIdx := strings.Index(tag, ",")
	if commaIdx == -1 { // 没有逗号
		return tag, OtherTagOptions("")
	}
	return tag[:commaIdx], OtherTagOptions(tag[commaIdx+1:])
}

// 测试直接传入结构体的情况
func parseOriginalStruct(v *url.Values, s any) error {
	// 使用reflect 取出s的值和类型
	var (
		sValue = reflect.ValueOf(s)
		sType  = reflect.TypeOf(s)
	)
	// 解引用指针和接口类型，直到得到实际的结构体类型
	sKind := sType.Kind()
	//fmt.Printf("【Original before】 is => %v, sType is => %v，sKind is %v\n", sValue, sType, sType.Kind())
	// 如果是指针或者接口类型 获取其指向的实际的值和类型
	if sKind == reflect.Ptr || sKind == reflect.Interface {
		sValue = sValue.Elem()
		sType = sValue.Type()
	}
	//如果不是结构体类型的值 就报错
	if sType.Kind() != reflect.Struct {
		return errors.New("can't parse s")
	}
	//fmt.Printf("【Original after】 is => %v, sType is => %v，sKind is %v\n", sValue, sType, sType.Kind())
	// 遍历结构体中的所有字段 sType.NumField() 返回的是结构体中的field数量
	for i := 0; i < sType.NumField(); i++ {
		var tagName string             // 最终解析出的field的tag key
		curField := sValue.Field(i)    // 获取当前field的值
		curFieldType := sType.Field(i) // 获取当前field的类型
		// 如果当前字段不可接口化（也就是不可导出，即无法被外部访问），则跳过
		if !curField.CanInterface() {
			continue
		}
		// 获取field中的json tag
		jsonTag := curFieldType.Tag.Get("json")
		fmt.Printf("curField is %#v，curFieldType is %v，是否可接口化 %t， jsonTag => %#v， curField value is %#v\n", curField, curFieldType, curField.CanInterface(), jsonTag, curField.Interface())
		// 如果json tag 为-，表示要忽略该字段，就不做处理
		if jsonTag == "-" {
			continue
		}
		// 解析tag
		tagName, tagOptions := parseTag(jsonTag)
		// 如果json tag 为空 默认转为小写
		if jsonTag == "" {
			tagName = strings.ToLower(curFieldType.Name)
		}
		var (
			omitEmptyValue bool
		)
		omitEmptyValue = strings.Contains(string(tagOptions), "omitempty")
		fmt.Printf("tagName is %s, options is %s\n", tagName, tagOptions)
		// 获取当前field的value值
		curFieldValue := fmt.Sprintf("%v", curField.Interface())
		if !(omitEmptyValue && len(curFieldValue) == 0) {
			v.Add(tagName, curFieldValue)
		}
	}
	return nil
}

// 测试传入结构体指针的情况
func parsePointerStruct(v *url.Values, s any) error {
	// 使用reflect 取出s的值和类型
	var (
		sValue = reflect.ValueOf(s)
		sType  = reflect.TypeOf(s)
	)
	// 解引用指针和接口类型，直到得到实际的结构体类型
	sKind := sType.Kind()
	//fmt.Printf("【Pointer before】sValue is => %v, sType is => %v，sKind is %v\n", sValue, sType, sType.Kind())
	// 如果是指针或者接口类型 获取其指向的实际的值和类型
	if sKind == reflect.Ptr || sKind == reflect.Interface {
		sValue = sValue.Elem()
		sType = sValue.Type()
	}
	//如果不是结构体类型的值 就报错
	if sType.Kind() != reflect.Struct {
		return errors.New("can't parse s")
	}
	//fmt.Printf("【Pointer after】sValue is => %v, sType is => %v，sKind is %v\n", sValue, sType, sType.Kind())
	return nil
}

func main() {
	var user = User{
		UserName: "张三",
		Gender:   1,
	}
	var u = url.Values{}
	err := parseOriginalStruct(&u, user)
	if err != nil {
		fmt.Println("解析出错")
	}
	fmt.Printf("解析后的参数u为 %#v\n", u)
	//parsePointerStruct(&user)
}
