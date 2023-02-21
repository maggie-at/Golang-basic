package basic

import (
	"fmt"
	"time"
)

// go语言中的循环控制只有for和for range, 没有while
func Loops() {
	// 1. for / for ... := range xxx {}
	i := 1
	for i <= 3 {
		fmt.Print(i)
		i++
	}
	fmt.Println("————————————————————————————————————————————————")

	for j := 0; j <= 9; j++ {
		fmt.Print(j)
	}
	fmt.Println("————————————————————————————————————————————————")

	slice := [...]int{1, 3, 5, 7, 9}
	for idx, v := range slice {
		fmt.Println("for range", idx, v)
	}
	for {
		fmt.Print("Loop forever")
		break
	}
	fmt.Println("————————————————————————————————————————————————")


	// 2. if-else if-else
	for n := 0; n <= 5; n++ {
		// 注意: 在 Go 中, 条件语句的圆括号不是必需的, 但是花括号是必需的
		// Go 没有三目运算符, 即使是基本的条件判断, 依然需要使用完整的if语句
		if n%3 == 0 {
			fmt.Println(".")
		} else if n%3 == 1 {
			fmt.Println("*")
		} else {
			fmt.Println("@")
		}
	}
	fmt.Println("————————————————————————————————————————————————")


	// 3. switch, case直接不需要break分隔, 默认只会执行一个case, 如果要执行多个需要使用fallthrough关键字
	j := 2

	// 3.1 switch 变量 {...}
	// 这种情况下, 如果前面的case没有break, 还是会进入default
	switch j {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Default")
	}

	// 同一个case中, 可以用逗号分隔多个表达式
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend :)")
	default:
		fmt.Println("Weekday :(")
	}

	// 3.2 switch {...} 不带表达式的switch ==> 等价于if/else逻辑
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// 类型开关 type switch, 比较类型而非值, 可以用来发现一个接口值的类型
	whatType := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a boolean")
		case int:
			fmt.Println("I'm an integer")
		default:
			fmt.Printf("%T is unknown\n", t)
		}
	}
	whatType(true)
	whatType("hello")
	whatType(100)
}
