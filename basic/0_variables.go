package basic

import "fmt"

func getAnonymousVar() (int, string, int){
	return 1, "Maggie", 25
}
func Variables_() {
/* 	// 变量声明
	var name string = "Maggie"
	var age int = 25
	fmt.Println(name, "-", age) */

/* 	// 批量声明
	var (
		name string = "Maggie"
		age int = 25
	)
	fmt.Println(name, "-", age) */

/* 	// 类型推断
	var name = "Maggie"
	var age = 25
	fmt.Printf("name: %v  age: %v", name, age) */

/* 	// 短变量声明 (只能在内部使用)
	name := "Maggie"
	age := 25
	fmt.Printf("name: %v  age: %v", name, age) */

	// 匿名变量 (函数式)
	_, name, age := getAnonymousVar()
	fmt.Printf("name: %v  age: %v", name, age)

}
