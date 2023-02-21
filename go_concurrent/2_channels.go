package go_concurrent

import (
	"fmt"
	"math/rand"
	"time"
)

var ch1 = make(chan int)

func send() {
	rand.Seed(time.Now().UnixNano()) // 设置随机种子
	sentVal := rand.Intn(10)         // 生成[0, n)之间的随机数
	fmt.Println(sentVal)
	time.Sleep(time.Second * 3) // 假装等待...
	ch1 <- sentVal
}
func Channels_() {
	defer close(ch1) // 函数返回前关闭通道
	go send()
	fmt.Println("Waiting...")
	receivedVal := <-ch1
	fmt.Println("Received:", receivedVal)
	fmt.Println("Closed")
}
