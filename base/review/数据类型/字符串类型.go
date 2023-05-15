package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func main() {

	str := `1234` + `sss`
	fmt.Println(str, "len：", len(str))
	//	获取ACSII
	strlen := len(str)
	fmt.Println("str的len：", strlen)
	str2 := "12,张三" + "123"
	//	获取unicode编码
	str2len := utf8.RuneCountInString(str2)
	fmt.Println(str2len)

	//	高性能拼接方式
	s1 := "hello"
	s2 := "world"
	var strinBuilder bytes.Buffer
	strinBuilder.WriteString(s1)
	strinBuilder.WriteString(s2)
	fmt.Println(strinBuilder.String())

	//获取Unicode码对应字符
	s4 := "hello 测试"
	fmt.Println("s4:", string([]rune(s4)[6]))
}
