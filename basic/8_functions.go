package basic

import (
	"fmt"
)

// 0. 值传递
func modify(a int) { // 不会改变a的值, 因为这里使用的是拷贝值
	a = a + 1
}
func modifySlice(s []int) { // 会改变原切片中的元素值
	s[0] = s[0] + 1
}

// 1. 普通函数
func add3(a, b, c int) int {
	return a + b + c
}

// 2. 多返回值
func multiRes() (int, int, int) {
	return 1, 3, 5
}

// 3. 可变参数
func sum(nums ...int) int {
	fmt.Println(nums)
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// 4. 闭包
// 解释: intSeq函数返回一个在其函数体内定义的匿名函数, 返回的函数使用闭包的方式「隐藏」变量i, 返回的函数「隐藏」变量i以形成闭包
// func intSeq() (func() int) {...}
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// 5. 递归
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
func Functions_() {

	// 0. 值传递
	aa := 1
	modify(aa)
	fmt.Println(aa) // 1

	ss := []int{1, 2, 3}
	modifySlice(ss)
	fmt.Println(ss) // [2 2 3]

	// 1. 普通函数
	res := add3(1, 3, 5)
	fmt.Println(res)
	fmt.Println("————————————————————————————————————————————————")

	// 2. 多返回值
	a, b, c := multiRes()
	fmt.Println(a, b, c)
	fmt.Println("————————————————————————————————————————————————")

	// 3. 变参函数
	// 3.1 调用方式1: 可变数量参数
	total := sum(1, 3, 5)
	fmt.Println(total)

	// 3.2 调用方式2: 切片
	nums := []int{1, 3, 5}
	total = sum(nums...)
	fmt.Println(total)
	fmt.Println("————————————————————————————————————————————————")

	// 4. 闭包
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println("————————————————————————————————————————————————")

	// 重新创建一个函数可以发现, 状态对于一个特定的函数是唯一的
	nextInt = intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// 5. 递归
	fmt.Println(fact(7))

	// 闭包也可以是递归的, 但这要求在定义闭包之前用类型化的 var 显式声明闭包
	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(7))
}
