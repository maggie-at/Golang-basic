package basic

import "fmt"

// 开闭原则(OCP): 对扩展开放, 对修改关闭

type Pet interface {
	eat()
	sleep()
}

// 定义了两个结构体Dog和Cat实现了Pet接口
type Dog struct {
	name string
}
type Cat struct {
	name string
}

func (d Dog) eat() {
	fmt.Printf("Dog %v is eating...\n", d.name)
}
func (d Dog) sleep() {
	fmt.Printf("Dog %v is sleeping...\n", d.name)
}
func (c Cat) eat() {
	fmt.Printf("Cat %v is eating...\n", c.name)
}
func (c Cat) sleep() {
	fmt.Printf("Cat %v is sleeping...\n", c.name)
}

type Host struct {
	name string
}

func (h Host) care(pet Pet) {
	pet.eat()
	pet.sleep()
}

func OCP_() {
	h := Host{
		name: "Jimmy",
	}
	dog := Dog{
		name: "Puppy",
	}
	cat := Cat{
		name: "Kitty",
	}
	h.care(dog)
	h.care(cat)
}
