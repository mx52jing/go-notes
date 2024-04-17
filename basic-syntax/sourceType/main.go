package main

import (
	"fmt"
	"math"
	"unsafe"
)

// 有符号整型
func SignedInteger() {
	var int8Val int8 = math.MaxInt8
	var int16Val int16 = math.MaxInt16
	var int32Val int32 = math.MaxInt32
	var int64Val int64 = math.MaxInt64
	var intVal int = math.MaxInt
	fmt.Printf("int8Val的类型是 %T, int8Val的大小 %d, int8Val的值是 %d\n", int8Val, unsafe.Sizeof(int8Val), int8Val)
	fmt.Printf("int16Val的类型是 %T, int16Val的大小 %d, int16Val的值是 %d\n", int16Val, unsafe.Sizeof(int16Val), int16Val)
	fmt.Printf("int32Val的类型是 %T, int32Val的大小 %d, int32Val的值是 %d\n", int32Val, unsafe.Sizeof(int32Val), int32Val)
	fmt.Printf("int64Val的类型是 %T, int64Val的大小 %d, int64Val的值是 %d\n", int64Val, unsafe.Sizeof(int64Val), int64Val)
	fmt.Printf("intVal的类型是 %T, intVal的大小 %d, intVal的值是 %d\n", intVal, unsafe.Sizeof(intVal), intVal)
}

// 无符号整型
func UnsignedInteger() {
	var uint8Val uint8 = math.MaxUint8
	var uint16Val uint16 = math.MaxUint16
	var uint32Val uint32 = math.MaxUint32
	var uint64Val uint64 = math.MaxUint64
	var uintVal uint = math.MaxUint
	fmt.Printf("uint8Val的类型是 %T, uint8Val的大小 %d, int8Val的值是 %d\n", uint8Val, unsafe.Sizeof(uint8Val), uint8Val)
	fmt.Printf("uint16Val的类型是 %T, uint16Val的大小 %d, int16Val的值是 %d\n", uint16Val, unsafe.Sizeof(uint16Val), uint16Val)
	fmt.Printf("uint32Val的类型是 %T, uint32Val的大小 %d, int32Val的值是 %d\n", uint32Val, unsafe.Sizeof(uint32Val), uint32Val)
	fmt.Printf("uint64Val的类型是 %T, uint64Val的大小 %d, int64Val的值是 %d\n", uint64Val, unsafe.Sizeof(uint64Val), uint64Val)
	fmt.Printf("uintVal的类型是 %T, uintVal的大小 %d, intVal的值是 %d\n", uintVal, unsafe.Sizeof(uintVal), uintVal)
}

func testFloat() {
	var float32Val float32 = math.MaxFloat32
	var float64Val float64 = math.MaxFloat64
	fmt.Printf("float32Val的类型是%T, float32Val是%g\n", float32Val, float32Val)
	fmt.Printf("float32V64的类型是%T, float64Val是%g\n", float64Val, float64Val)
}

func testBool() {
	var a bool = true
	b := false
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	fmt.Println("true && false = ", a && b)
	fmt.Println("true || false = ", a || b)
}

func main() {
	// SignedInteger()
	// UnsignedInteger()
	// testFloat()
	testBool()
}