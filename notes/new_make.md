### new 🆚 make

> - 使用范围: 
>   - `make`用来分配及初始化类型为`slice`, `map`, `chan`的数据
>   - `new`可以分配任意类型的数据
> - 返回值:
>   - `make`返回引用(结构体本身), 即`T`
>   - `new`接收类型作为参数, 返回一个指向该类型的指针, 即`*T`
> - 空间:
>   - `make`分配后会进行初始化
>   - `new`分配的空间是空的

```GO
slice := make([]int, 0, 100)    // struct {data, cap, len}
hash := make(map[int]bool, 10)  // 指向runtime.hmap结构体的指针
ch := make(chan int, 5)         // 指向runtime.hchan结构体的指针
```

