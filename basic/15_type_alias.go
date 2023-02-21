package basic

import "fmt"

func TypeAlias_() {
	// 新类型, 结构体也是这样定义的
	type MyInt int
	var i MyInt
	fmt.Printf("%T\n", i) // basic.MyInt

	// 类型别名
	type Str = string
	var s Str
	fmt.Printf("%T\n", s) // string
}
