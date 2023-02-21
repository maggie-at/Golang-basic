package basic

import "fmt"

func Slice_() {
	// slice的类型「仅由它所包含的元素的类型决定」, 「与元素个数无关」
	// Data uintptr, Len int, Cap int  ->  Cap>=Len

	// 定义方式一: 类似数组, 但不指定长度
	t := []string{"A", "C", "E"}
	fmt.Println("definition 2:", t)

	// 定义方式二: make([]type, initial_len)
	s := make([]string, 3)
	fmt.Println("definition 1:", s)

	//定义方式三: part of existed array / slice
	u := s[0:2]
	fmt.Println("definition 3:", u)

	// 二维切片
	// Slice可以组成多维数据结构, 「内部的slice长度可以不一致」, 这一点和多维数组不同
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		twoD[i] = make([]int, i+1) // 内部slice长度可以不一致
		for j := 0; j < i+1; j++ {
			twoD[i][j] = j
		}
	}
	fmt.Println("2d slice:", twoD)

	// set
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)

	// get
	fmt.Println("get:", s[2])

	// len
	fmt.Println("len: ", len(s))

	// 内建函数 append: 返回一个包含一个或者多个新值的slice
	s = append(s, "d")
	s = append(s, "e", "f", "g")
	fmt.Println(s)

	// copy: 创建一个空的与s相同(或者更大)长度的slice, 将s复制给destS
	destS := make([]string, len(s)+5)
	copy(destS, s)
	fmt.Println(destS)

	// 切片: slice[low: high], 左闭右开
	part1, part2, part3 := s[:3], s[3:6], s[6:]
	fmt.Println("Part:", part1, part2, part3)

	// 不超过容量(与len不同)则空值填充
	fmt.Println("len: ", len(s))
	longer := s[:10]
	fmt.Println("longer slice", longer)
}
