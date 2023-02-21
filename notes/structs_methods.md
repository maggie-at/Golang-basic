### struct

> GO语言没有面向对象的概念, 可以用结构体来实现继承、组合等特性

```GO
type person struct {
    id   int
	name string
	age  int
}
```

### methods

> 方法: 可以为「值类型」或者「指针类型」的「接收者(receiver)」定义方法


> 「为结构体类型定义方法(methods)」, 使用`receiver.方法()`来调用
> 
> 值类型的结构体作为接收参数, 调用方法时会对结构体进行拷贝, 无法修改结构体的值

```GO
// 值传递
func (r rect) calcArea() int {
	r.name = "area11111"		// 不会对外部结构体产生修改
	return r.width * r.height
}
// 引用传递: 指针类型作为接收参数, 可以避免调用方法时产生一个拷贝, 可以对结构体值进行修改
func (r *rect) calcPerim() int {
	// 实际是(*r).name, 这里是自动解引用
	r.name = "perim22222"		// 会对外部结构体产生修改
	return 2 * (r.width + r.height)
}

// 调用
r.calcArea()
r.calcPerim()
```
