package go_concurrent

import (
	"fmt"
	"runtime"
)

func showGoSched() {
	for i := 0; i < 5; i++ {
		fmt.Println("show()")
	}
}
func showGoexit() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		if i >= 5 {
			runtime.Goexit() // runtime.Goexit(): 退出当前协程
		}
	}
}
func Runtime_() {

	go showGoSched()
	// 1. runtime.Gosched()
	// Gosched yields the processor, allowing other goroutines to run.
	// It does not suspend the current goroutine, so execution resumes automatically.
	// 让当前goroutine让出CPU, 好让其它的goroutine获得执行的机会; 同时，当前的goroutine也会在未来的某个时间点继续运行。
	runtime.Gosched()
	fmt.Println("main()...")

	/*
		// 2. runtime.Goexit()
		// Calling Goexit from the main goroutine terminates that goroutine
		go showGoexit()
		time.Sleep(time.Second * 1) // 防止Runtime_()方法退出导致观察不到showGoexit()中的现象
	*/

	/*
		// 3. runtime.GOMAXPROCS(i)
		fmt.Printf("runtime.NumCPU(): %v\n", runtime.NumCPU())
		runtime.GOMAXPROCS(2)
	*/
}
