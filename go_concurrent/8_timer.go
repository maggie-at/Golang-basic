package go_concurrent

import (
	"fmt"
	"time"
)

func Timer_() {
	// 实例化1: timer.NewTimer(duration)
	timer1 := time.NewTimer(time.Second * 3)

	fmt.Println("begin: ", time.Now())

	t := <-timer1.C // C: 从这个channel读, 一直会在这里阻塞, 直到计时器设定的时间结束
	fmt.Println("timer.C: ", t)

	// 实例化2: time.After()
	<-time.After(time.Second * 3)
	fmt.Println("After:", time.Now())

	// 协程间timer使用
	timer2 := time.NewTimer(time.Second * 3)
	go func() {
		<-timer2.C
		fmt.Println("anonymous func")
	}()
	time.Sleep(time.Second * 5) // 让主协程等一下, 不然匿名函数没执行完就退出了

	// timer.Stop()
	fmt.Println(time.Now())
	timer3 := time.NewTimer(time.Second * 3)
	stop := timer3.Stop()
	if stop {
		fmt.Println("stop: ", stop)
	}

	// timer.Reset(duration)
	timer4 := time.NewTimer(time.Second * 3)
	timer4.Reset(time.Second * 10)
	<-timer4.C // 阻塞等待计时器结束
	fmt.Println("Reset")
}
