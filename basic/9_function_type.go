package basic

import "fmt"

// 定义函数类型
type compareFT func(int, int) int

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func compare(rule compareFT, a int, b int) int {
	return rule(a, b)
}
func FunctionType_() {
	var ft_min compareFT = min
	var ft_max compareFT = max

	fmt.Printf("%T", ft_max) // basic.compareFT1

	fmt.Println(ft_min(1, 2))
	fmt.Println(ft_max(1, 2))

	// compareFT可以作为函数参数, 有点像「多态」
	// https://blog.csdn.net/gfdsgsfdsfds/article/details/117264622
	minVal := compare(min, 1, 2)
	maxVal := compare(max, 1, 2)
	fmt.Println(minVal)
	fmt.Println(maxVal)
}
