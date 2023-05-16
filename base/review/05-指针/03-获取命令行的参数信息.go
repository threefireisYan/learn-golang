package main

import (
	"flag"
	"fmt"
)

// 定义命令行
var cmdMode = flag.String("mode:", "", "运行模式，可以设置为fast")

func main() {
	//	解析命令行
	flag.Parse()
	fmt.Println("运行模式为：", *cmdMode)
}
