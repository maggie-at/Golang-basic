package go_packages

import (
	"fmt"
	"sort"
)

// UIntSlice 自定义对任意类型的排序, 实现sort接口
type UIntSlice []uint

func (u UIntSlice) Len() int {
	return len(u)
}
func (u UIntSlice) Less(i, j int) bool {
	return u[i] < u[j]
}
func (u UIntSlice) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

// TwoDimSlice 自定义对二维数组的排序, 实现sort接口
type TwoDimSlice [][]int

func (tds TwoDimSlice) Len() int {
	return len(tds)
}
func (tds TwoDimSlice) Less(i, j int) bool {
	return tds[i][1] < tds[j][1]
}
func (tds TwoDimSlice) Swap(i, j int) {
	tds[i], tds[j] = tds[j], tds[i]
}

// StuInfo 根据自定义结构体的某一个字段进行排序, 同样实现三个sort接口的方法
type StuInfo struct {
	Name string
	Age  int
}
type StuList []StuInfo

func (stu StuList) Len() int {
	return len(stu)
}
func (stu StuList) Less(i, j int) bool {
	return stu[i].Age < stu[j].Age
}
func (stu StuList) Swap(i, j int) {
	stu[i], stu[j] = stu[j], stu[i]
}
func Sort_() {
	s := []int{2, 4, 1, 3}
	sort.Ints(s)
	fmt.Println(s)

	// 自定义对「任意类型」的排序
	us := UIntSlice{2, 4, 1, 3}
	sort.Sort(us)
	fmt.Println(us)

	// 对「二维切片」的任意一维排序
	tds := TwoDimSlice{{1, 9}, {2, 7}, {3, 8}}
	sort.Sort(tds)
	fmt.Println(tds)

	// 对「自定义结构体」排序
	stu := StuList{
		{Name: "Alan", Age: 25},
		{Name: "Maggie", Age: 24},
	}
	sort.Sort(stu)
	fmt.Println(stu)
}
