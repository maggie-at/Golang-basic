### Standard Library

> https://pkg.go.dev/

#### `os`
> https://pkg.go.dev/os


#### `log`
> https://pkg.go.dev/log
> 
> - `Print`: 打印日志
> - `Panic`: 打印日志, 抛出panic异常, 顺序代码不会执行, **`defer`会执行**
> - `Fatal`: 打印日志, 强制程序执行(`os.Exit(1)`), **`defer`不会执行**

```GO
logger := log.New(f, "Prefix: ", log.Ldate|log.Ltime|log.Lshortfile)
```

#### `builtin`
> - `append(slice []Type, elems ...Type) []Type`
> - `panic`: 顺序代码不会执行, **`defer`会执行**


#### `bytes`
> - `bytes.Contains(b, subslice []byte) bool`
> - `bytes.Count(s, sep []byte) int`
> - `bytes.Repeat(b []byte, count int) []byte`
> - `bytes.Replace(s, old, new []byte, n int) []byte` 
> - `bytes.Runes(s []byte) []rune`


#### `error`
> - 创建一个`error` => `errors.New("some description")`
> - 自定义`XxxError`结构体, 实现`error`接口的`Error() string`方法, 即实现了`error`接口


#### `sort`
> 自定义排序需要实现下面的接口中的三个方法

```GO
// An implementation of Interface can be sorted by the routines in this package.
type Interface interface {
	Len() int            // 返回集合元素个数 
	Less(i, j int) bool  // i>j, 返回索引i位置的元素是否小于索引j位置的元素
	Swap(i, j int)       // 交换索引i和索引j位置的值
}
```


#### `time`
> https://pkg.go.dev/time


#### `json`
> https://pkg.go.dev/json

```GO
type JsonNode struct {
	K1 string `json:"k1"`
	K2 string `json:"k2"`
	K3 int    `json:"k3"`
}
func JSON_() {
	var jsonStr string = `{"k1": "v111","k2": "v222","k3": 333}`
	fmt.Println(jsonStr)

	// 1. 存入map
	// GO中可以用空接口interface{}表示Object
	var mp map[string]interface{}
	// Unmarshal: 先将字符串转为[]byte数组, 然后传入map指针
	json.Unmarshal([]byte(jsonStr), &mp)
	fmt.Println(mp)

	// 2. 直接存入struct, 需要提前定义好struct
	var node JsonNode
	json.Unmarshal([]byte(jsonStr), &node)
	fmt.Println(node)

	// 3. io流
	// 从json文件中读取
	f, _ := os.Open("go_packages/test.json")
	defer f.Close()
	decoder := json.NewDecoder(f)
	var mp1 map[string]interface{}
	decoder.Decode(&mp1)
	fmt.Println(mp1)
	
    // 写入json文件
    var newJson string = `{"a1": "b111","a2": "b222","a3": 333}`
    wf, _ := os.OpenFile("test.json", os.O_WRONLY, 0777)
    defer wf.Close()
    encoder := json.NewEncoder(wf)
    encoder.Encode(newJson)
}
```


#### `math`
> https://pkg.go.dev/math
> 
> https://pkg.go.dev/math/rand