package go_concurrent

import (
	"fmt"
	"sync"
)

// WaitGroup实现协程同步, 保证主函数在所有协程结束后退出
var wg_ sync.WaitGroup
var i int = 100

// 互斥锁(加在代码段前后)
var lock sync.Mutex // 加Mutex并不能保证先后顺序, 而是add()/sub()过程互斥

func add() {
	lock.Lock()      // 加锁
	defer wg_.Done() // wg--
	i += 1
	fmt.Println("i++:", i)
	lock.Unlock() // 解锁
}
func sub() {
	lock.Lock()      // 加锁
	defer wg_.Done() // wg--
	i -= 1
	fmt.Println("i--:", i)
	lock.Unlock() // 解锁
}
func Mutex_() {
	/*
		// 使用go修饰, 每次运行结果不保证正确, 每次运行结果也可能相同
		for i := 0; i < 100; i++ {
			go add()
			go sub()
		}
	*/

	for i := 0; i < 100; i++ {
		wg_.Add(1) // wg++
		go add()
		wg_.Add(1) // wg++
		go sub()
	}
	// Wait blocks until the WaitGroup counter is zero.
	wg_.Wait() // WaitGroup实现协程同步, 保证主函数在所有协程结束后退出
}
