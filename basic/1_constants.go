package basic

import "fmt"

func Constants_() {
	// const必须在定义时就进行初始化
	const PI1 float32 = 3.14
	const PI2 = 3.1415926
	fmt.Printf("%v\n%v\n", PI1, PI2)

	const (
		PI_bit1 = "3"
		PI_bit2 = "."
		PI_bit3 = "1"
		PI_bit4 = "4"
	)
	fmt.Printf("%v%v%v%v\n", PI_bit1, PI_bit2, PI_bit3, PI_bit4)

	// iota: 每调用一次加1, 遇到const关键字被重置
	const (
		a1 = iota
		a2 = iota
		a3 = iota
	)
	fmt.Println(a1) // 0
	fmt.Println(a2) // 1
	fmt.Println(a3) // 2

	const a4 = iota
	fmt.Println(a4) // 0

	const (
		A1  = iota
		_   // +1跳过
		A2  = iota
		tmp = 111 // +1跳过
		A3  = iota
	)
	fmt.Println(A1) // 0
	fmt.Println(A2) // 2
	fmt.Println(A3) // 0
}
