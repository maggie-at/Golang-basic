package basic

import "fmt"

func Array_() {
	// 可以将数组类型记成 [len]type , 数组长度不可变
	// 1. 定义: 创建一个刚好可以存放 5 个 int 元素的数组 a, 元素会设置为类型默认值
	var a [5]int
	fmt.Println("define", a)

	// set: 使用 array[index] = value 来设置数组指定位置的值
	a[3] = 10
	fmt.Println("set:", a)
	// get: 使用 array[index] 得到值
	fmt.Println("get:", a[3])
	// len
	fmt.Println("len", len(a))

	// 2. 简略定义
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	// 3. 隐式定义数组长度
	var c = [...]int{1, 2, 3, 4, 5}
	fmt.Println(c)
	
	// 4. 使用索引值初始化
	var d = [5]int{0: 1, 2: 4}
	fmt.Println(d)

	// range遍历
	for idx, v := range d{
		fmt.Printf("%v: %v\n", idx, v)
	}

	// 二维数组
	var twoD [3][4]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}
