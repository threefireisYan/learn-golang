package main

import "fmt"

func main() {
	str := "12345678测试"
	//	1.普通for循环遍历字符串(不支持中文)
	for i := 0; i < len(str); i++ {
		fmt.Println("str单字符:", str[i])
	}
	//	2.for range遍历（支持中文,输出时为Unicode码）
	for key, value := range str {
		fmt.Println("key:", key, ",str单字符:", str[key], ",value:", value)
	}

}
