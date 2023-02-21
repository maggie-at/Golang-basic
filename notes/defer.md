### defer

> GO中的`defer`语句会将其修饰的语句进行**延迟处理**, 在`defer`所在的函数即将返回时(正常顺序执行的语句全部执行完), 将`defer`修饰的语句按照**逆序**(先进后出)执行。
> 
> 也就是说, 函数返回前, 先被`defer`的语句最后被执行, 最后被`defer`的语句先被执行

> `defer`特性
> - `defer`用于注册延迟调用
> - 这些调用直到`return`前才被执行, 因此可以用来做**资源清理**
> - 多个`defer`语句按照「先进后出」的方式执行
> - `defer`语句中的变量, 在`defer`声明时就决定了

> `defer`用途
> - 关闭文件句柄
> - 锁资源释放
> - 数据库连接释放

```GO
// 输出顺序: 0 4 3 2 1

line := 0
fmt.Println("start:", line)

// defer语句中的变量, 在defer声明时就决定了
line += 1
defer fmt.Printf("line: %v\n", line)

line += 1
defer fmt.Printf("line: %v\n", line)

line += 1
defer fmt.Printf("line: %v\n", line)

line += 1
fmt.Println("end:", line)
```