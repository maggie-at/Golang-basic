### GO指针

> GO语言中的函数传参是**值传递**, 如果想通过函数修改某个变量, 可以传递指向该变量地址的指针, 而无序拷贝数据
>
> GO语言中, 指针类型不能进行偏移和运算, 只有`&(取地址)`和`*(根据地址取值)`两个操作

```GO
var int_p *int
fmt.Printf("%T: %v", int_p, int_p)
```