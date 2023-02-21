package basic

import "fmt"

type person struct {
	name string
	age  int
}

// 构造方法(自己定义的, 不是由GO提供的, 作用是初始化并返回一个结构体指针)
// 返回值: 可以安全地返回指向局部变量的指针, 因为局部变量将在函数的作用域中继续存在
func newPerson(name string) *person {
	p := person{name: "Alan"}
	p.age = 25
	return &p
}
func Structs_() {
	// 可以部分初始化
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 40})

	// 在构造函数中封装创建新的结构实例
	fmt.Println(newPerson("Jon"))

	s := person{
		name: "Sean",
		age:  24,
	}
	s.age = 25
	fmt.Println(s.age)

	// 也可以对结构体指针使用"." - 指针会被自动解引用
	sp := &s
	sp.age = 26
	fmt.Println(sp.age)
}
