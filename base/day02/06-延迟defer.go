package main

import (
	"fmt"
	"os"
)

func main() {
	//	1.延迟，关键字，可以用于修饰语句，函数，确保这条语句可以在当前栈推出的时候执行
	//	2.一般用做资源清理工作
	//	3.解锁、关闭文件、关闭数据库连接等
	//	4.在同一个函数中多次调用defer，执行时类似于栈的机制，先入后出
	readFile("01-switch.go")
}

func readFile(filename string) {
	//1.go语言一般会将错误码作为最后一个参数返回
	//2.err一般nil代表没有错误，执行成功，非nil标识执行失败
	f1, err := os.Open(filename)
	defer func(a int) {
		fmt.Println("准备关闭文件,code", a)
		f1.Close()
	}(100) //匿名函数 加上（）就是需要运行
	if err != nil {
		fmt.Println("打开文件失败,os.Open,err:", err)
		return
	}
	buf := make([]byte, 1024)
	n, _ := f1.Read(buf)
	fmt.Println("读取数据长度：", n)
	fmt.Println("读取的内容：", string(buf))

}
