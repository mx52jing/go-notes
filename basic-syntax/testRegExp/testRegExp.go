package main

import (
	"fmt"
	"regexp"
)


var (
	dateRegExp = regexp.MustCompile(`9999[-_/]?12[-_/]?31`)
)

func main() {
	str := "9999_12_31"
	fmt.Println(dateRegExp.FindAllStringSubmatch(str, -1))
	fmt.Printf("dateRegExp.FindAllStringSubmatchIndex(str, -1): %v\n", dateRegExp.FindAllStringSubmatchIndex(str, -1))
}