package basic

import "fmt"
import "sort"

func Map_() {
	// 定义一: 直接声明
	mp1 := map[string]int{
		"foo": 1, 
		"bar": 2, 
	}
	fmt.Println(mp1)

	// 定义二: make声明
	mp2 := make(map[string]int)
	mp2["k0"] = 2
	mp2["k1"] = 4
	mp2["k2"] = 8
	fmt.Println(mp2)
	
	// 定义三: 声明并直接初始化
	mp3 := map[string]int {
		"a": 97,
		"b": 98,
		"c": 99,
	}

	// 内建函数len: 返回map中键值对数目
	fmt.Println(len(mp3))

	// 内建函数delete: 用key索引, 从map中移除键值对
	delete(mp2, "k0")

	// 第二个返回值: 表明map中是否存在这个键
	_, ok := mp2["k0"]
	fmt.Println("exist:", ok)

	_, ok2 := mp2["k1"]
	fmt.Println("exist:", ok2)

	// range遍历 (无序)
	for k, v := range mp3{
		fmt.Println(k, "->", v)
	}

	// 面试题: 如何对map进行有序排序
	// 可以先把`key`收集到一个`slice`里, 给`slice`排序, 再循环输出对应的`mp[key]`即可
	mp := make(map[int]string)
	mp[3] = "three"
	mp[1] = "one"
	mp[2] = "two"

	slice := make([]int, 0)
	for k := range mp{
		slice = append(slice, k)
	}

	sort.Ints(slice)

	for _, ele := range slice{
		fmt.Println(ele, "->", mp[ele])
	}
}
