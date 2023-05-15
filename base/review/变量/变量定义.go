package main

import (
	"fmt"
	"net"
)

func main() {
	//	1.使用var定义变量数据类型
	var name string
	name = "张三"
	var num int32
	num = 123
	var float float32
	float = 123.22
	var test string = "1234"
	fmt.Println("name:", name, ",num:", num, ",float:", float, ",test:", test)

	//	2.批量声明
	var (
		name2 int
		str   string
		flag  bool
	)
	name2 = 12
	str = "123"
	flag = false
	fmt.Println("name2:", name2)
	fmt.Println("str:", str)
	fmt.Println("flag", flag)

	//	不指定变量数据类型定义
	i := 1
	i2 := "测试"
	i3 := true

	fmt.Println(i, i2, i3)

	//	多重赋值
	var conn net.Conn
	var err error
	conn, err = net.Dial("tcp", "127.0.0.1:8080")

	fmt.Println(conn)
	fmt.Println(err)

}
