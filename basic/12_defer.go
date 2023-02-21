package basic

import "fmt"

func Defer_ () {
	// 输出顺序: 0 4 3 2 1

	line := 0
	fmt.Println("start:", line)
	
	// defer语句中的变量, 在defer声明时就决定了
	line += 1
	defer fmt.Printf("line: %v\n", line)

	line += 1
	defer fmt.Printf("line: %v\n", line)

	line += 1
	defer fmt.Printf("line: %v\n", line)

	line += 1
	fmt.Println("end:", line)
}