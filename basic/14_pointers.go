package basic

import "fmt"

// 值传递, 得到实参的拷贝
func zeroVal(ival int) {
	ival = 0
}
// 引用传递, 得到实参的指针
func zeroPtr(iptr *int) {
	*iptr = 2
}
func Pointers_() {
	var i int = 1
	var iPtr *int = &i
	fmt.Printf("%T: %v\n", iPtr, iPtr)

	zeroVal(i)
	fmt.Println("zeroVal:", i)

	// 通过 &i 语法来取得i的内存地址, 即指向i的指针
	zeroPtr(&i)
	fmt.Println("zeroPtr:", i)


	// 指针数组
	var arr = [3]int{100, 200, 300}
	var ptrs [3]*int
	for i:=0; i<3; i++ {
		ptrs[i] = &arr[i]
	}
	for i=0; i<3; i++ {
		fmt.Printf("arr[%d] = %d\n", i, *ptrs[i])
	}
}