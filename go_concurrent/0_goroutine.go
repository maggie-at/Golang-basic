package go_concurrent

import (
	"fmt"
)

func Child() {
	fmt.Println("goroutine")
	for {
		fmt.Println("live...")
	}
}
func Parent() {
	go Child() // 如果Parent中没有sleep, 并没有打印Child中的内容 => 主死从随 => 解决方法之一: main协程sleep

	//time.Sleep(time.Millisecond * 10)

	fmt.Println("main end")
}
