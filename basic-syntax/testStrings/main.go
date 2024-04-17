package main

import (
	"fmt"
	"strings"
)

func main() {
	// 判断字符串str是否以prefix 开头
	// strings.HasPrefix(str, prefix)
	// str1 := "hash是开发计划"
	// fmt.Println(strings.HasPrefix(str1, "hash")) // true
	// fmt.Println(strings.HasPrefix(str1, "ha")) // true
	// fmt.Println(strings.HasPrefix(str1, "hashs")) // false

	// // 判断字符串str是否以sufffix 结尾
	// // strings.HasSuffix(str, sufffix)
	// str2 := "asd是否会words"
	// fmt.Println(strings.HasSuffix(str2, "s")) // true
	// fmt.Println(strings.HasSuffix(str2, "wo")) // false
	// fmt.Println(strings.HasSuffix(str2, "会words")) // true

	// // 判断字符串包含关系
	// str3 := "whatgowotk"
	// fmt.Println(strings.Contains(str3, "go")) // true
	// fmt.Println(strings.Contains(str3, "og")) // false
	// fmt.Println(strings.Contains(str3, "wotk")) // true

	// str5 := "shfwuebrqr"
	// fmt.Println(strings.Index(str5, "ud")) // -1
	// fmt.Println(strings.Index(str5, "fwu")) // 2
	// fmt.Println(strings.Index(str5, "r")) // 7


	// str6 := "shfwuebrqr"
	// fmt.Println(strings.Index(str6, "r")) // 7
	// fmt.Println(strings.Index(str6, "she")) // -1
	// fmt.Println(strings.Index(str6, "s")) // 0

	// str7 := "abgcuyabiucu"
	// fmt.Println(strings.Replace(str7, "ab", "--", 1)) // --gcuyabiucu
	// fmt.Println(strings.Replace(str7, "ab", "--", 2)) // --gcuy--iucu
	// fmt.Println(strings.Replace(str7, "u", "=", -1)) // abgc=yabi=c=

	// str8 := "abgcuyabiucudd"
	// fmt.Println(strings.Count(str8, "a")) // 2
	// fmt.Println(strings.Count(str8, "ab")) // 2
	// fmt.Println(strings.Count(str8, "u"))  // 3
	// fmt.Println(strings.Count(str8, "d")) // 2
	// fmt.Println(strings.Count(str8, "dd")) // 1

	// fmt.Println(strings.Repeat("a", 2)) // aa
	// fmt.Println(strings.Repeat("a-b", 2)) // a-ba-b
	// fmt.Println(strings.Repeat("a b", 2)) // a ba b

	// str9 := "UhtKWlwTsO"
	// fmt.Println(strings.ToLower(str9)) // uhtkwlwtso
	// fmt.Println(strings.ToUpper(str9)) // UHTKWLWTSO

	// str10 := " ehe ll o world e"
	// str11 := "hell o world eh"
	// fmt.Println(strings.TrimSpace(str10)) // "ehe ll o world e"
	// fmt.Println(strings.Trim(str11, "h")) // "ell o world e"
	// fmt.Println(strings.TrimLeft(str11, "e")) // "hell o world eh"
	// fmt.Println(strings.TrimRight(str11, "eh")) // "hell o world"

	// str12 := "hello world"
	// str13 := "张,李,孙,"
	// str14 := ""
	// fmt.Printf("值为 %v, 类型为 %T\n", strings.Fields(str12), strings.Fields(str12))  // 值为 [hello world], 类型为 []string
	// fmt.Printf("值为 %v, 类型为 %T\n", strings.Fields(str14), strings.Fields(str14)) // 值为 [], 类型为 []string
	// fmt.Printf("值为 %v, 类型为 %T\n", strings.Split(str13, ","), strings.Split(str13, ",")) // 值为 [张 李 孙 ], 类型为 []string

	name := []string{"张环", "李朗", "沈韬", "肖豹"}
	fmt.Println(strings.Join(name, "-")) // "张环-李朗-沈韬-肖豹"
	fmt.Println(strings.Join(name, " ")) // "张环 李朗 沈韬 肖豹"
	fmt.Println(strings.Join(name, "%")) // "张环%李朗%沈韬%肖豹"
}