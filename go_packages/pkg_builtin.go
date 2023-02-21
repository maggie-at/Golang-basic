package go_packages

import "fmt"

func Builtin_() {
	// append: 追加
	// 切片的元素追加和扩容: https://draveness.me/golang/docs/part2-foundation/ch03-datastructure/golang-array-and-slice/#324-%E8%BF%BD%E5%8A%A0%E5%92%8C%E6%89%A9%E5%AE%B9
	s := []int{0, 25, 50}
	s = append(s, 100, 200)
	fmt.Println(s)

	s2 := []int{400, 800}
	s = append(s, s2...) // ...用来打散slice, 转化成参数列表
	fmt.Println(s)
}

func New_Make() {
	// new可以分配任意类型的数据, 返回值类型为指针 *T
	b := new(bool)
	fmt.Printf("%T: %v\n", b, *b)

	i := new(int)
	fmt.Printf("%T: %v\n", i, *i)

	s := new(string)
	fmt.Printf("%T: len = %v\n", s, len(*s))

	// new([]int)返回一个新分配的、被置零的slice结构体的指针, 即指向 nil 的slice的指针
	var s1 *[]int = new([]int)
	fmt.Println(*s1, " len: ", len(*s1), " cap: ", cap(*s1)) // []  len:  0  cap:  0

	// make([]int, len, cap)
	var s2 []int = make([]int, 10, 20)
	fmt.Println(s2, " len: ", len(s2), " cap: ", cap(s2)) // [0 0 0 0 0 0 0 0 0 0]  len:  10  cap:  20

}
