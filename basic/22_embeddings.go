package basic

import "fmt"

// 和「21_inherit.go」一个意思
// Go支持对于结构体(struct)和接口(interfaces)的嵌入(embedding), 以表达一种更加无缝的组合(composition)类型

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// 一个container嵌入了一个base, 一个嵌入看起来像一个没有名字的字段
type container struct {
	base
	str string
}

func embedding_func() {
	c := container{
		base: base{num: 11},
		str:  "string...",
	}
	fmt.Println(c.base.num)
	fmt.Println(c.num)
	fmt.Printf("co={num: %v, str: %v}\n", c.num, c.str)

	fmt.Println(c.describe())
	fmt.Println(c.base.describe())
}
