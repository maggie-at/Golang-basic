### Ticker
> Timer只执行一次, Ticker可以「周期性执行」

```GO
// Ticker containing a channel that will send the current time on the channel after each tick.
ticker := time.NewTicker(time.Second)

counter := 1
for t := range ticker.C {
    fmt.Println(t)
    counter++
    if counter > 5 {
        break
    }
}
```

```GO
// 周期性发送 / 接收
chanInt := make(chan int)
// 用一个协程写入chan
go func() {
    for _ = range ticker.C {
		// select多个case满足, 随机执行一个
        select {
        case chanInt <- 1:
            fmt.Println("Send 1")
        case chanInt <- 2:
            fmt.Println("Send 2")
        case chanInt <- 3:
            fmt.Println("Send 3")
        }
    }
}()

sum := 0
for v := range chanInt {
    fmt.Println("Received ", v)
    sum += v
    if sum >= 10 {
        break
    }
}
```