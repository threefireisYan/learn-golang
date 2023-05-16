package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	//定义一个包含Unicode的字符串 替换JAVA为GO
	var str1 string = "hello world 学习JAVA使我快乐!"
	var str string = "GO"
	//使用for循环检查每个字符所在下标
	for key, value := range str1 {
		fmt.Println("下标:", key, ",下标对应值:", value)
	}
	//获取JAVA的初始下标
	index1 := strings.Index(str1, "JAVA")
	fmt.Println("初始下标:", index1)
	//截取下标前的位置
	str2 := str1[:index1]
	fmt.Println("初始下标前字符串：", str2)
	//	截取下标+4后的字符串 (包含当前下标字符:不包含当前下标字符)
	str3 := str1[index1+4:]
	fmt.Println(str3)
	//	拼接最后的字符串
	var lastStr bytes.Buffer
	lastStr.WriteString(str2)
	lastStr.WriteString(str)
	lastStr.WriteString(str3)
	fmt.Println(lastStr.String())

}
