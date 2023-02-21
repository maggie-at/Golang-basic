package basic

import (
	"fmt"
	"math"
)

func hello() {
	// 1. hello-world
	fmt.Println("hello world")
	fmt.Println("————————————————————————————————————————————————")

	// 2. values
	fmt.Println("go" + "lang")

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println("————————————————————————————————————————————————")

	// 3. variables
	// 3.1 var
	var s1 string
	var s2 = "complete initial"
	fmt.Println(len(s1), s2)

	var b1 bool
	var b2 bool = true
	fmt.Println(b1, b2)

	var i1 int
	var i2 int = 3
	fmt.Println(i1, i2)

	// 3.2 :=
	s3 := "short initial"
	fmt.Println(s3)

	// 3.3 var a,b,... =
	var a, b, c int = 1, 2, 4
	fmt.Println(a, b, c)
	fmt.Println("————————————————————————————————————————————————")

	// 4. constant
	// 常数表达式可以执行任意精度的运算
	// 数值型常量没有确定的类型, 直到被给定某个类型, 比如显式类型转化
	// 一个数字可以根据上下文的需要（比如变量赋值、函数调用）自动确定类型
	const n = 500000000 // Cannot (re)assign a new value to n
	const d = 3e20 / n
	fmt.Println("d =", d)
	fmt.Println("int64(d) =", int64(d))
	fmt.Println("sin(d) =", math.Sin(n))
	// 举个例子, 这里的math.Sin函数需要一个float64的参数, n会自动确定类型

	fmt.Println("————————————————————————————————————————————————")
}
