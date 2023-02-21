package basic

import "fmt"

// 执行顺序: 变量初始化 -> init() -> main()
var i int = initVariable()

func initVariable() int {
	fmt.Println("————initVariable()————")
	return 100
}

// init()不需要调用, 导入包时就会自动执行
// 每个包可以有多个init(), 执行顺序没有明确规定, 因此定义时不要依赖顺序
func init() {
	fmt.Println("————init()————")
}
