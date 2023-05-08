package main

import (
	"fmt"
	"os"
)

// 从命令行输入参数，在switch中进行处理
func main() {

	// Go: os.Args ==>直接可以获取命令输入，是一个字符串切片

	cmds := os.Args
	//os.Args[0] =>程序名字
	//os.Args[1] =>第一个参数
	for key, value := range cmds {
		fmt.Println("key:", key, ",value:", value, ",len", len(cmds))
	}

	if len(cmds) < 2 {
		fmt.Println("请正确输入参数")
		return
	}
	//在Go语言中默认有break
	//如果想向下穿透，那么需要关键字：fallthrough
	switch cmds[1] {
	case "hello":
		fmt.Println("ok")
		fallthrough
	case "world":
		fmt.Println("ok2")
	default:
		fmt.Println("other")
	}

	fmt.Println("END")
}
