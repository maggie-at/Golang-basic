package go_concurrent

import "fmt"

func ChannelIteration_() {
	ch := make(chan int) // 一个unbuffered通道

	go func() {
		for i := 0; i < 2; i++ {
			ch <- i
		}
		close(ch) // (如果这个是在main中) 如果通道没有close并且读请求多, 会出现死锁 => fatal error: all goroutines are asleep - deadlock!
	}()

	// 避免死锁的方式1: close(chan)
	for i := 0; i < 3; i++ {
		r := <-ch
		fmt.Println("read ", r) // 如果通道close了, 会返回该类型默认是, int型通道会返回0
	}

	// 避免死锁的方式2: v, ok := <-ch, 用ok判断chan为空
	for {
		v, ok := <-ch
		if ok {
			fmt.Println(v)
		} else {
			break
		}
	}
}
