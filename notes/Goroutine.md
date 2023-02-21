### Go-routine

#### WaitGroup - 协程同步
> `wg.Add(i int)`: 一般在创建Goroutine前Add, 将计数器`+i`
> 
> `wg.Done()`: 等价于`wg.Add(-1)`, 将计数器`+1`
> 
> `wg.Wait()`: 等待直到WaitGroup的计数器为0

```GO
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
```


#### runtime
##### 1. `runtime.Gosched()`
> 当前Goroutine让出CPU, 让其它Goroutine获得执行的机会
> 
> 同时, 当前的Goroutine也会在未来的某个时间点继续运行

```GO
// 打印顺序是show()*5, 然后才会打印main()
func showGoSched() {
    for i := 0; i < 5; i++ {
        fmt.Println("show()")
    }
}
func main(){
    go showGoSched()
    // 1. Runtime.Gosched()
    // Gosched yields the processor, allowing other goroutines to run.
    // It does not suspend the current goroutine, so execution resumes automatically.
    runtime.Gosched()
    fmt.Println("main()...") // 打印main这句会在最后执行
}
```

##### 2. `Goexit()`
> 退出当前协程

```GO
func showGoexit() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		if i >= 5 {
			runtime.Goexit() // 退出当前协程
		}
	}
}
func Runtime_() {
    // 2. Runtime.Goexit()
    go showGoexit()
    time.Sleep(time.Second * 1) // 防止Runtime_()方法退出导致观察不到showGoexit()中的现象
}
```