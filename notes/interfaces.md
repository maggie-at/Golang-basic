### interface

> 在Go中, 一个类只需要实现了接口要求的所有函数, 就说这个类实现了该接口
>
> - 一个类型可以实现多个接口
> 
> - 多个类型也可以实现同一个接口 (多态)

```GO
// 接口定义
type geometry interface {
	getArea() float64
	getPerim() float64
}
```

```GO
// 定义准备实现geometry接口的结构体
type rectangle struct {
	width, height float64
}
type circle struct {
	radius float64
}
```

```GO
// 接口的实现
// rectangle实现geometry接口定义的方法
func (r rectangle) getArea() float64 {
	return r.width * r.height
}
func (r rectangle) getPerim() float64 {
	return 2*r.width + 2*r.height
}
// circle实现geometry接口定义的方法
func (c circle) getArea() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) getPerim() float64 {
	return 2 * math.Pi * c.radius
}
```

```GO
// 接口的使用: 这里是一个以geometry接口作为参数的方法, 可以传入任何实现了该接口所有方法的结构体
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.getArea())
	fmt.Println(g.getPerim())
}
func interfaces_func() {
	r := Rectangle{width: 5, height: 3}
	c := Circle{radius: 3}

	measure(r)
	measure(c)
}
```

```GO
// 多态
var rGeo geometry
rGeo = Rectangle{width: 5, height: 3}
fmt.Printf("rGeo.getArea(): %v\n", rGeo.getArea())

var cGeo geometry
cGeo = Circle{radius: 3}
fmt.Printf("cGeo.getArea(): %v\n", cGeo.getArea())
```