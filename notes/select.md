### select

> `select`是GO中一个控制结构, 用于处理异步IO操作, 类似于`switch`, **`select`会监听`case`语句中`channel`的读写操作**
> 
> 当`case`中`channel`读写操作为非阻塞状态时(可以读写的状态), 将会触发相应动作
> 
> `select`中的`case`语句必须是一个`channel`操作, `select`中的`default`语句总是可运行的

> 运行机制:
> 
> - 如果多个`case`都可以执行, `select`会随机公平地选出一个执行, 其它不会执行
> 
> - 如果没有可运行的`case`语句, 且有`default`语句, 那么就会执行`default`的操作
> 
> - 如果没有可运行的`case`语句, 且没有`default`语句, `select`将阻塞, 直到某个`case`满足条件可以运行

```GO
var chanInt = make(chan int)
var chanStr = make(chan string)

func Select_() {
	go func() {
		chanInt <- 11
		chanStr <- "Maggie"
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
```