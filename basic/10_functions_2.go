package basic

import "fmt"

func add(a int, b int) int {
	return a + b
}
func minus(a int, b int) int {
	return a - b
}
func multiply(a int, b int) int {
	return a * b
}
func cal(op string) func(int, int) int {
	switch op {
	case "+":
		return add
	case "-":
		return minus
	case "*":
		return multiply
	default:
		return nil
	}
}
func Functions__() {
	addFunc := cal("+")
	minusFunc := cal("-")
	multiplyFunc := cal("*")

	// func(int, int) int
	fmt.Printf("%T\n%T\n%T\n", addFunc, minusFunc, multiplyFunc)

	fmt.Printf("addFunc(2, 1): %v\n", addFunc(2, 1))
	fmt.Printf("minusFunc(2, 1): %v\n", minusFunc(2, 1))
	fmt.Printf("multiplyFunc(2, 1): %v\n", multiplyFunc(2, 1))

	// 匿名函数
	max := func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	fmt.Printf("max(1, 2): %v\n", max(1, 2))
}
