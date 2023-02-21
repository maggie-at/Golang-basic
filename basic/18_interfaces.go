package basic

import (
	"fmt"
	"math"
)

// 接口
type geometry interface {
	getArea() float64
	getPerim() float64
}

// 结构体
type Rectangle struct {
	width, height float64
}
type Circle struct {
	radius float64
}

// 实现接口的方法
func (r Rectangle) getArea() float64 {
	return r.width * r.height
}
func (r Rectangle) getPerim() float64 {
	return 2*r.width + 2*r.height
}
func (c Circle) getArea() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) getPerim() float64 {
	return 2 * math.Pi * c.radius
}

// 接口的使用: 这里是一个以geometry接口作为参数的方法, 可以传入任何实现了该接口所有方法的结构体
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.getArea())
	fmt.Println(g.getPerim())
}
func Interface_() {
	r := Rectangle{width: 5, height: 3}
	c := Circle{radius: 3}

	measure(r)
	measure(c)

	// 多态
	var rGeo geometry
	rGeo = Rectangle{width: 5, height: 3}
	fmt.Printf("rGeo.getArea(): %v\n", rGeo.getArea())

	var cGeo geometry
	cGeo = Circle{radius: 3}
	fmt.Printf("cGeo.getArea(): %v\n", cGeo.getArea())

}
