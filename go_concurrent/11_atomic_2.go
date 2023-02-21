package go_concurrent

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

func AtomicOperations_() {
	// 原子操作: 数值运算
	var n int32 = 100
	for i := 0; i < 100; i++ {
		go atomic.AddInt32(&n, 2)
		go atomic.AddInt32(&n, -2)
	}
	fmt.Println(i)

	// 原子操作: 读 / 写
	go func() {
		for {
			atomic.StoreInt32(&n, int32(rand.Intn(10)))
			<-time.After(time.Second * 1)

		}
	}()
	go func() {
		for {
			fmt.Println(atomic.LoadInt32(&n))
			<-time.After(time.Second * 1)
		}
	}()
	time.Sleep(time.Second * 5)
}
