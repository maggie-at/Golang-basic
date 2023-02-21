### Timer
> `Timer`就是定时器, 可以实现定时操作, 内部也是通过`channel`实现的

```GO
// The Timer type represents a single event.
// When the Timer expires, the current time will be sent on C, unless the Timer was created by AfterFunc.
// A Timer must be created with NewTimer or AfterFunc.
type Timer struct {
	C <-chan Time
	r runtimeTimer
}
```

```GO
// 创建Timer: A Timer must be created with NewTimer or AfterFunc.
// 实例化1: timer.NewTimer(duration)
timer1 := time.NewTimer(time.Second * 3)

fmt.Println("begin: ", time.Now())

t := <-timer1.C // C: 从这个channel读, 一直会在这里阻塞, 直到计时器设定的时间结束
fmt.Println("timer.C: ", t)
```

```GO
// 实例化2: time.After()
<-time.After(time.Second * 3)
fmt.Println("After:", time.Now())
```

```GO
// 协程间timer使用
timer2 := time.NewTimer(time.Second * 3)
go func() {
    <-timer2.C
    fmt.Println("anonymous func")
}()
time.Sleep(time.Second * 5) // 让主协程等一下, 不然匿名函数没执行完就退出了
```

```GO
// timer.Stop()
fmt.Println(time.Now())
timer3 := time.NewTimer(time.Second * 3)
stop := timer3.Stop()
if stop {
    fmt.Println("stop: ", stop)
}
```

```GO
// timer.Reset(duration)
timer4 := time.NewTimer(time.Second * 3)
timer4.Reset(time.Second * 10)
<-timer4.C // 阻塞等待计时器结束
fmt.Println("Reset")
``` 