package basic

import "fmt"

type rect struct {
	name   string
	width  int
	height int
}

// 方法: 可以为「值类型」或者「指针类型」的「接收者」定义方法 => 将函数和结构体绑定后达到OOP的效果

// 「为结构体类型定义方法(methods)」, 类似"面向对象"中「对象.方法()」的感觉
// 值类型作为接收参数, 调用方法时会对结构体进行拷贝, 并且无法修改结构体的值
func (r rect) area() int {
	r.name = "area11111"
	return r.width * r.height
}

// 指针类型作为接收参数, 可以避免调用方法时产生一个拷贝, 并且可以对结构体值进行修改
func (r *rect) perim() int {
	r.name = "perim22222"
	return 2 * (r.width + r.height)
}
func Methods_() {
	r := rect{width: 5, height: 3}
	fmt.Println(r)

	fmt.Println("area:", r.area())
	fmt.Println(r)

	fmt.Println("perim", r.perim())
	fmt.Println(r)
}
