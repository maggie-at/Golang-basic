package basic

import (
	"fmt"
	"strings"
)

func counter() func(int) int {
	x := 0
	return func(y int) int{
		x += y
		return x
	}
}
func makeSuffix(suffix string) func(string) string {
	return func(file string) string {
		if !strings.HasSuffix(file, suffix) {
			return file + suffix
		}
		return file
	}
}
// 举例3: 两个返回值(函数), 比如计算账户余额
func calculate(base int) (func(int) int, func(int) int) {
	// 两个匿名函数
	acc := func(x int) int {
		base += x
		return base
	}
	sub := func(x int) int {
		base -= x
		return base
	}
	return acc, sub
}
func Closure_(){
	// 变量adder是一个函数, 它引用了外部作用域的x变量, 此时adder就是一个闭包
	// 在闭包adder的生命周期内, 变量x一直有效
	// 发生了变量逃逸, 从栈中跑到堆中了
	adder := counter()
	fmt.Printf("adder(1): %v\n", adder(1))	// 1
	fmt.Printf("adder(1): %v\n", adder(1))	// 2

	// 重新调用counter是不同的生命周期
	adder = counter()
	fmt.Printf("adder(1): %v\n", adder(1))	// 1
	fmt.Printf("adder(1): %v\n", adder(1))	// 2


	// 举例2: 加后缀
	jpgHelper := makeSuffix(".jpg")
	txtHelper := makeSuffix(".txt")
	fmt.Printf("jpgHelper(\"aaa\"): %v\n", jpgHelper("aaa"))
	fmt.Printf("jpgHelper(\"bbb\"): %v\n", jpgHelper("bbb"))
	fmt.Printf("txtHelper(\"bbb\"): %v\n", txtHelper("bbb"))
	fmt.Printf("txtHelper(\"ccc\"): %v\n", txtHelper("ccc"))


	// 举例3: 加减计算器
	acc, sub := calculate(0)

	fmt.Printf("acc(800): %v\n", acc(800))
	fmt.Printf("acc(1000): %v\n", acc(1000))
	fmt.Printf("acc(300): %v\n", acc(300))
	fmt.Printf("sub(2000): %v\n", sub(2000))
}