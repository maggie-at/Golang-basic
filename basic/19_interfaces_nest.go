package basic

import "fmt"

type Flyer interface{
	fly()
}
type Swimmer interface{
	swim()
}
type Superman interface{
	Flyer
	Swimmer
}
type People struct{
	name string
}
func (s People) fly() {
	fmt.Println(s.name + " can fly")
}
func (s People) swim() {
	fmt.Println(s.name + " can swim")
}
func Interface_Nest_(){
	var p Superman	// 向上类型转换
	p = People{
		name: "Herry",
	}
	p.fly()
	p.swim()
}