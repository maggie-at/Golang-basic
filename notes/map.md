### map

> map是无序的(不按输入顺序, 且不按key排序)
> 
> 面试题: 如何对map进行有序排序
> 
> 可以先把`key`收集到一个`slice`里, 给`slice`排序, 再循环输出对应的`mp[key]`即可

```GO
mp := make(map[int]string)
mp[3] = "three"
mp[1] = "one"
mp[2] = "two"

slice := make([]int)
for k, _ := range mp{
    slice = append(slice, k)
}

sort.Ints(slice)

for _, ele := range slice{
    fmt.Println(ele, "->", mp[ele])
}
```

```GO
// 定义一: 直接声明
var a map[string]int
a["a"] = 97
a["b"] = 98
```

```GO
// 定义二: make声明
b := make(map[string]int)
b["A"] = 65
b["B"] = 66
```

```GO
// 定义三: 声明并直接初始化
c := map[string]int {
    "a": 97,
    "b": 98,
}
```

```GO
// key是否存在, mp[k]第二维表示是否存在
v, ok := mp[k]
```

```GO
// range遍历 (无序)
for k, v := range mp3{
    fmt.Println(k, "->", v)
}
```