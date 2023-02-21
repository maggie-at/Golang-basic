### function

> **返回值**: 可以有0~多个返回值, 返回值需要执行类型, 可以没有名称

```GO
// 1. 无返回值
func add3(a, b, c int)  {
	fmt.Println(a + b + c)
}

// 2. 多返回值
func multiRes() (int, int, int) {
	return 1, 3, 5
}
```

> **参数**: 值传递, 声明函数时的参数列表为形参, 调用函数时的参数列表为实参

```GO
func modify(a int) {		// 不会改变a的值, 因为这里使用的是拷贝值
	a = a+1
}
// 切片, 数组, map是可以指向同一块内存的
func modifySlice(s []int) {	// 会改变原切片中的元素值
	s[0] = s[0] + 1
}
```

```GO
// 可变参数
func sum(nums ...int) int {
	fmt.Println(nums)
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
// 可以传入任意数量参数或者切片
// 调用方式1: 可变数量参数
total := sum(1, 3, 5)
fmt.Println(total)

// 调用方式2: 切片
nums := []int{1, 3, 5}
total = sum(nums...)
fmt.Println(total)
```

```GO
// 匿名函数
max := func (a int, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}
```