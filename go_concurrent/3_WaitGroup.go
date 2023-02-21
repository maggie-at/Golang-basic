package go_concurrent

import (
	"fmt"
	"sync"
)

// WaitGroup实现协程同步
var wg sync.WaitGroup // 一个正在执行协程的计数器

func hello(i int) {
	defer wg.Done() // 或者是wg.Add(-1), 记录一个任务的完成
	fmt.Println(i)
}

func WaitGroup_() {
	for i := 0; i < 10; i++ {
		wg.Add(1) // 记录一个任务的启动
		go hello(i)
	}
	// Wait blocks until the WaitGroup counter is zero.
	wg.Wait()
	fmt.Println("main end")
}
