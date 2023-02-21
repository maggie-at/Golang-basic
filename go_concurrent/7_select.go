package go_concurrent

import (
	"fmt"
	"time"
)

var chanInt = make(chan int)
var chanStr = make(chan string)

func Select_() {
	go func() {
		chanInt <- 25
		chanStr <- "Alan"
		defer close(chanInt)
		defer close(chanStr)
	}()

	for {
		select {
		// 如果执行了close, 仍然可以合法读到chan的内容, 只不过是默认值
		case r := <-chanInt:
			fmt.Println("chanInt: ", r)
		case r := <-chanStr:
			fmt.Println("chanStr: ", r)
		// 如果没有close, 也没有default, 就会出现死锁
		default:
			fmt.Println("Default") // 如果没有close(chan), 就会执行default
		}
		time.Sleep(time.Second * 1)
	}
}
