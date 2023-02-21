package basic

import (
	"fmt"
	"strings"
	"unicode/utf8"
	"bytes"
)

func Strings_() {
	// Go语言中的字符串是「以UTF-8为编码段文本容器」, 等价于[]byte
	const s = "一段字符串"
	fmt.Println(s)
	fmt.Println(len(s))

	// ``定义的字符串允许换行
	const longStr = `
	line1
	line2
	line3
	`
	fmt.Printf("?%v", longStr)


	// 字符串的组成
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])		// %x: 十六进制
	}

	// 在Go语言当中, 字符的概念被称为 rune - 它是一个表示 Unicode 编码的整数
	fmt.Println("\nRune count:", utf8.RuneCountInString(s))

	// range 循环专门处理字符串并解码每个 rune 及其在字符串中的偏移量
	for idx, runeVal := range s {
		fmt.Printf("%#U starts at %d\n", runeVal, idx)
	}


	// 字符串拼接
	s1 := "abc"
	s2 := "xyz"
	// 方式一: +
	fmt.Println(s1+s2)

	// 方式二: res = fmt.Sprintf("%s ... %s", s1, s2)
	s3 := fmt.Sprintf("%s ... %s", s1, s2)
	fmt.Println(s3)

	// 方式三: strings.Join()
	s4 := strings.Join([]string{s1, s2}, " ~ ")
	fmt.Println(s4)

	// 方式四: buffer.WriteString()
	var buffer bytes.Buffer
	buffer.WriteString(s1)
	buffer.WriteString(" ~ ")
	buffer.WriteString(s2)
	fmt.Println(buffer.String())


	// 字符串切片 (左开右闭)
	str := "Hello string"
	fmt.Println(str[1])
	fmt.Println(str[1:])
	fmt.Println(str[1:4])
	fmt.Println(str[:4])
}
