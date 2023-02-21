package basic

import (
	"fmt"
	"unsafe"
	"math"
)
func tmpFunc() string {
	return "A function"
}
func Types() {
	// 整型
	var _int int = 100
	var _int2 = 100
	_int3 := 100
	fmt.Printf("%T  %v\n", _int, _int)
	fmt.Printf("%T  %v\n", _int2, _int2)
	fmt.Printf("%T  %v\n", _int3, _int3)

	var i8 int8
	var i32 int32
	var ui8 uint8
	var ui32 uint32

	// int8  1B  -128~127
	fmt.Printf("%T  %dB  %v~%v\n", i8, unsafe.Sizeof(i8), math.MinInt8, math.MaxInt8)
	// int32  4B  -2147483648~2147483647
	fmt.Printf("%T  %dB  %v~%v\n", i32, unsafe.Sizeof(i32), math.MinInt32, math.MaxInt32)
	// uint8  1B  max: 255
	fmt.Printf("%T  %dB  max: %v\n", ui8, unsafe.Sizeof(ui8), math.MaxUint8)
	// uint32  4B  max: 4294967295
	fmt.Printf("%T  %dB  max: %v\n", ui32, unsafe.Sizeof(ui32), math.MaxUint32)

	// 浮点型
	var f32 float32 = 0.32
	var f64 float64 = 0.64
	fmt.Printf("%.2f\n", f32)
	fmt.Printf("%.6f\n", f64)
	// float32  4B max: 3.4028234663852886e+38
	fmt.Printf("%T  %dB max: %v\n", f32, unsafe.Sizeof(f32), math.MaxFloat32)
	// float64  8B max: 1.7976931348623157e+308
	fmt.Printf("%T  %dB max: %v\n", f64, unsafe.Sizeof(f64), math.MaxFloat64)

	// 字符串
	var _str string = "aaa"
	var _str2 = "bbb"
	_str3 := "ccc"
	fmt.Printf("%T  %v\n", _str, _str)
	fmt.Printf("%T  %v\n", _str2, _str2)
	fmt.Printf("%T  %v\n", _str3, _str3)

	// 指针类型
	a := 100
	p := &a
	fmt.Printf("%T\n", p)

	// 数组类型: [5]int
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%T\n", arr)

	// 切片类型: []int
	slice := []int{1,2,3}
	fmt.Printf("%T\n", slice)

	// 函数类型
	fmt.Printf("%T\n", tmpFunc)	// func() string

}