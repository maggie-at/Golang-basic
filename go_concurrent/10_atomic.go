package go_concurrent

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 解决方案二: atomic(CAS)
var m int32 = 100

func addAtomic() {
	atomic.AddInt32(&m, 1)
}
func subAtomic() {
	atomic.AddInt32(&m, -1)
}

// 解决方式一: 互斥锁sync.Mutex
var lock_n sync.Mutex

var n int = 100

func addWithLock() {
	lock_n.Lock()
	n++
	lock_n.Unlock()
}
func subWithLock() {
	lock_n.Lock()
	n--
	lock_n.Unlock()
}

func Atomic_() {
	// 解决方案一: sync.Mutex
	for i := 0; i < 100; i++ {
		go addWithLock()
		go subWithLock()
	}
	time.Sleep(time.Second * 3)
	fmt.Println(n)

	// 解决方案二: atomic
	for i := 0; i < 100; i++ {
		go addAtomic()
		go subAtomic()
	}
	time.Sleep(time.Second * 3)
	fmt.Println(m)
}
