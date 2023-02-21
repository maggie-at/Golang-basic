package go_packages

import (
	"bytes"
	"fmt"
)

// Type_Trans 类型转换
func Type_Trans() {
	var i int = 100
	var b byte
	b = byte(i)
	fmt.Println(b)
	i = int(b)
	fmt.Println(i)

	// 字符串 <-> []byte字节切片
	var str string = "你好, world"

	bs := []byte(str)
	fmt.Println(bs)

	str = string(bs)
	fmt.Println(str)
}
func Bytes_() {
	bs := []byte("你好, world lol")

	// bytes.Contains(b, subslice []byte) bool
	sub1 := []byte("好")
	sub2 := []byte("word")
	fmt.Printf("%v\n", bytes.Contains(bs, sub1)) // -1
	fmt.Printf("%v\n", bytes.Contains(bs, sub2)) // 1

	// bytes.Count(s, sep []byte) int
	sep1 := []byte("l")
	sep2 := []byte("o")
	fmt.Printf("%v\n", bytes.Count(bs, sep1))
	fmt.Printf("%v\n", bytes.Count(bs, sep2))

	// bytes.Repeat(b []byte, count int) []byte
	bs = []byte("ha")
	fmt.Println(string(bytes.Repeat(bs, 3)))

	// bytes.Replace(s, old, new []byte, n int) []byte
	bs = []byte("hello world")
	old := []byte("o")
	new := []byte("aaa")
	fmt.Println(string(bytes.Replace(bs, old, new, -1))) // -1表示不限次数
	fmt.Println(string(bs))

	// bytes.Join(s [][]byte, sep []byte) []byte
	bs_list := [][]byte{[]byte("你好"), []byte("世界")}
	sep := []byte(", ")
	fmt.Println(string(bytes.Join(bs_list, sep)))

	// bytes.Runes(s []byte) []rune
	// Runes interprets s as a sequence of UTF-8-encoded code points.
	// It returns a slice of runes (Unicode code points) equivalent to s.
	// 实现: make([]rune, utf8.RuneCount(s))
	bs = []byte("你好世界, lalala")
	r := bytes.Runes(bs)
	fmt.Println(len(bs))
	fmt.Println(len(r))
}
