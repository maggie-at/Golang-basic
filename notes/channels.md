### Channel 通道

#### 缓冲通道
> 缓冲通道有能力在接收到一个或多个值之前保存它们, 不强制`Goroutine`在同一时刻准备好执行发送和接收。
> 
> 当通道中没有要接收的值, 接收才会阻塞; 仅当没有可用缓冲区放置正在发送的值时, 发送才会阻塞

#### 无缓冲通道
> 无缓冲通道没有保存能力, 发送和接收`Goroutine`在任何发送或接收操作之前的同一时刻必须都准备就绪。
>
> 同步是在通道上发送和接收之间交互的基础。
> 
> 如果发送或者接收方没有在同一时刻准备好, 则通道会让执行各自发送或接收操作的`Goroutine`首先等待。

```GO
unbuffered := make(chan int)  // 整型无缓冲通道
buffered := make(chan int, 10)  // 整型有缓冲通道
```

#### 通道的发送和接收特性
> - 对于同一个通道, 发送操作之间是互斥的, 接收操作之间也是互斥的
> - 发送操作和接收操作中「对元素值的处理是不可分割的」
> - 发送操作 / 接收操作在完全完成前会被阻塞

```GO
var ch1 = make(chan int)

func send() {
	rand.Seed(time.Now().UnixNano()) // 设置随机种子
	sentVal := rand.Intn(10) // 生成[0, n)之间的随机数
	fmt.Println(sentVal)
	time.Sleep(time.Second * 3) // 假装等待...
	ch1 <- sentVal
}
func Channels_() {
	defer close(ch1) // 函数返回前关闭通道
	go send()
	fmt.Println("Waiting...")
	receivedVal := <-ch1 // 这条语句会等待直到通道ch1有数据才会执行
	fmt.Println("Received:", receivedVal)
	fmt.Println("Closed")
}
```