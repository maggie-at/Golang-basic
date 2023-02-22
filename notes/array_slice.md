### array

> 数组是由相同类型元素的集合组成的数据结构, 一旦定义**长度不可修改**, 可以通过下标访问
> 
> Go数组的长度是类型的一部分, 数组在初始化之后大小就无法改变了

> 数组的存储:
> - 当元素数量小于或者等于 4 个时, 会直接将数组中的元素放置在`栈`上
> - 当元素数量大于 4 个时，会将数组中的元素放置到`静态存储区`并在运行时取出
> - 无论是在栈上还是静态存储区, 数组在内存中都是一连串的内存空间, 我们通过指向数组开头的指针、元素的数量以及元素类型占的空间大小表示数组

> 数组的访问:
> - 编译时检查:
>   - 访问数组的索引是非整数时, 报错`"non-integer array index %v"`
>   - 访问数组的索引是负数时, 报错`"invalid array index %v (index must be non-negative)"`
>   - 访问数组的索引越界时, 报错`"invalid array index %v (out of bounds for %d-element array)"`
> - 运行时检查:
>   - Go 语言运行时在发现数组、切片和字符串的越界操作会由运行时的`runtime.panicIndex`和`runtime.goPanicIndex`触发程序的运行时错误并导致崩溃退出

```GO
// 方式一: var array_name [SIZE]variable_type
var arr [5]int
arr[0] = 1
arr[1] = 3
//  ...
```

```GO
// 方式二: 简短定义
b := [5]int{1,2,3,4,5}
```

```GO
// 方式三: 隐式指定数组长度, GO语言会在编译期间通过源代码推导数组的大小
// [...]T{1, 2, 3} 和 [3]T{1, 2, 3} 在运行时是完全等价的, [...]T 这种初始化方式只是GO的一种语法糖
c := [...]int{1, 2, 3, 4, 5}
```

```GO
// 方式四: 使用索引值初始化
d := [5]int{0: 1, 2: 4}
```

```GO
// range遍历
for idx, v := range arr{
    fmt.Printf("%v: %v\n", idx, v)
}
```


### slice

> 切片可以理解动态数组, 即**可变长度**的数组(长度不固定, 可以扩容), 有`len`和`cap`两个属性, 满足`cap>=len`
> 
> GO中切片类型的声明方式与数组有一些相似, 不过声明时只需要指定切片中的元素类型, 不需要指明长度: `[]int`, []interface{}`
> 
> 切片在编译期间的生成的类型只会包含切片中的元素类型，即`int`或者`interface{}`等
>
> 编译期间的切片是`cmd/compile/internal/types.Slice`类型的, 但是在运行时切片可以由`reflect.SliceHeader`结构体表示

```GO
type SliceHeader struct {
	Data uintptr    // 指向数组的指针;
	Len  int        // 切片长度
	Cap  int        // 切片容量, 等于Data数组的大小
}
```

> `(数组)Data`是一片连续的内存空间, 用于存储切片中的全部元素, 数组中的元素只是逻辑上的概念, 底层存储其实都是连续的, 所以可以将切片理解成「一片连续的内存空间」加上「长度」与「容量」的标识
> 
> ![SLICE struct](https://img.draveness.me/2019-02-20-golang-slice-struct.png)
> 
> 切片引入了一个抽象层, 提供了对数组中部分连续片段的引用, 而作为数组的引用, 我们可以在运行期间修改它的长度和范围
> 
> 当切片底层的数组长度不足时就会触发扩容, 切片指向的数组可能会发生变化, 不过在上层看来切片是没有变化的, 上层只需要与切片打交道而不需要关心数组的变化

> 通过两个冒号创建切片: `slice[x:y:z]`切片实体为`[x:y]`, 切片长度为`len = y-x`, 切片容量为`cap = z-x`
>
> ![通过两个冒号创建切片](/notes/assets/slice定义.png)


#### 初始化Slice
##### Slice创建方式一: 通过「下标」获得数组或切片的一部分
> ⚠️需要注意的是: 使用下标初始化切片「不会拷贝原数组或者原切片中的数据」, 它只会创建一个指向原数组的切片结构体
>
> 所以「下标方式创建切片时, 修改新切片的数据也会修改原切片」

```GO
u := s[0:2]
fmt.Println("definition 3:", u)
```

##### Slice创建方式二: 字面量
> 如果使用字面量的方式创建切片, 大部分的工作都会在编译期间完成

```GO
// []int{1, 2, 3} 创建新的切片的编译过程如下: 
var vstat [3]int    // 根据切片中的元素数量「对底层数组的大小进行推断并创建一个数组」
vstat[0] = 1
vstat[1] = 2        // 将这些字面量元素存储到初始化的数组中 
vstat[2] = 3
var vauto *[3]int = new([3]int)     // 创建一个同样指向 [len]T 类型的数组指针 
*vauto = vstat      // 将静态存储区的数组 vstat 赋值给 vauto 指针所在的地址
slice := vauto[:]   // 通过 [:] 操作获取一个底层使用 vauto 的切片
// 第5步中的 [:] 就是使用下标创建切片的方法, 从这一点我们也能看出 [:] 操作是创建切片最底层的一种方法
```

```GO
// 定义方式二: 字面量创建, 类似数组, 但不指定长度
t := []string{"A", "C", "E"}
fmt.Println("definition 2:", t)
```

##### Slice创建方式三: make
> 当我们使用`make`关键字创建切片时, 很多工作都需要运行时的参与
> - 调用方必须向`make`函数传入切片的大小以及可选的容量
> - 类型检查期间的`cmd/compile/internal/gc.typecheck1`函数会校验入参, 比如`len<=cap`

```GO
// 定义方式三: make([]type, initial_len, cap[optional])
s := make([]string, 3)
fmt.Println("definition 1:", s)
```


#### Slice追加元素
> 这里要注意的是，append函数执行完后，返回的是一个全新的 slice，并且对传入的 slice 并不影响。

```GO
// 添加元素 / 其它切片
s2 := append(s1, new_element)
s3 := append(s1, s2)
``

```GO
// 删除index位置的元素
s2 := append(s1[:index], s2[index:]...)
```

> 使用`append`向切片中追加元素是常见的切片操作, 中间代码生成阶段的`cmd/compile/internal/gc.state.append`方法会「根据返回值是否覆盖原变量」, 选择进入两种流程:
> 
> 是否覆盖原变量的逻辑其实差不多, 最大的区别在于得到的新切片是否会赋值回原变量
> 
> 如果选择覆盖原有的变量, 就不需要担心切片发生拷贝影响性能, 因为GO语言编译器已经对这种常见的情况做出了优化
> 
> ![append(slice, newElement)](https://img.draveness.me/2020-03-12-15839729948451-golang-slice-append.png)
```GO
// Case 1: 如果不需要覆盖回原变量
// newSlice := append(slice, 1, 2, 3)
ptr, len, cap := slice
newlen := len + 3
if newlen > cap {
    ptr, len, cap = growslice(slice, newlen) // 如果追加元素后切片大小>cap, 调用growslice扩容
    newlen = len + 3
}
*(ptr+len) = 1      // 加入新元素
*(ptr+len+1) = 2
*(ptr+len+2) = 3
return makeslice(ptr, newlen, cap)
```

```GO
// Case 2: 如果需要覆盖回原变量
// slice = append(slice, 1, 2, 3)
a := &slice
ptr, len, cap := slice
newlen := len + 3
if uint(newlen) > uint(cap) {
    newptr, len, newcap = growslice(slice, newlen)
    vardef(a)
    *a.cap = newcap
    *a.ptr = newptr
}
newlen = len + 3
*a.len = newlen
*(ptr+len) = 1
*(ptr+len+1) = 2
*(ptr+len+2) = 3
```


#### Slice扩容 - `runtime.growslice`

> 扩容是为切片分配新的内存空间并拷贝原切片中元素的过程
> 
> 分配内存空间前需要先确定新的切片容量, 运行时根据切片的当前容量选择不同的策略进行扩容:
> - 第一步: 确定大致容量
>   - 如果「期望容量大于当前容量的两倍」, 就会直接使用期望容量
>   - 如果「当前切片的长度小于 1024」就会将容量翻倍
>   - 如果「当前切片的长度大于 1024」就会每次增加`25%`的容量(1/4), 直到新容量满足期望容量
> - 第二步: 根据元素类型「对齐内存(`roundupsize()`)」, 减少内存碎片

```GO
func growslice(et *_type, old slice, cap int) slice {
	newcap := old.cap
	doublecap := newcap + newcap
	if cap > doublecap {
		newcap = cap
	} else {
		if old.len < 1024 {     // 注意这里是当前切片的「len」
			newcap = doublecap
		} else {
			for 0 < newcap && newcap < cap {
				newcap += newcap / 4    // 增加25%直到新容量>=期望容量
			}
			if newcap <= 0 {
				newcap = cap
			}
		}
	}
}
```


#### 浅拷贝和深拷贝

```GO
// 浅拷贝
var s1 = []int{1, 2, 3, 4, 5}
var s2 = s1
s1[2] = 333     // 此时两个切片指向同一块内存, 两者都会被修改
```

```GO
// 深拷贝
s3 := make([]int, 3)    // 要用make初始化空间
copy(s3, s1)    // copy(目的, 源)
```