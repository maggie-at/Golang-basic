package basic

import (
	"fmt"
)

// go的继承同样是通过嵌套(组合)实现的
type Animal struct {
	name string
	age  int
}

func (a Animal) sleep() {
	fmt.Printf("%v is sleeping\n", a.name)
}
func (a Animal) eat() {
	fmt.Printf("%v is eating\n", a.name)
}

type Dog_ struct {
	Animal        // 嵌套 => 继承
	skill  string // 对Animal的扩展
}
type Cat_ struct {
	Animal        // 嵌套 => 继承
	color  string // 对Animal的扩展
}

func Inherit_() {
	dog := Dog_{
		Animal{
			name: "Puppy",
			age:  2,
		},
		"bark", // 顺序初始化, 可以省略属性名
	}
	cat := Cat_{
		Animal{
			name: "Kitty",
			age:  3,
		},
		"black",
	}
	dog.sleep()
	cat.eat()
}
