package main

import (
	"fmt"
	"time"
)

func main() {
	//这里启用了一个协程
	go ph()
	time.Sleep(2 * time.Second)
	//需要主线程进行睡眠等待协程完成，否则主线程结束，未完成或正在运行的线程也会结束
	fmt.Println("你好世界")

}

func ph() {
	fmt.Println("hello world")
}
