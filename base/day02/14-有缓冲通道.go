package main

import (
	"fmt"
	"time"
)

func main() {

	//	1.有缓冲的通道,当缓冲写满后，写阻塞，当被读取后，再恢复写入
	//	2.当缓冲区读取完毕，读阻塞
	//	3.如果管道没有使用make分配空间，那么管道默认是nil的，读取、写入都会阻塞
	//	4.对于一个管道，读与写的次数一定要对等，否则会死锁

	//当管道的读写不一致的时候
	//如果阻塞在主go程，那么程序会崩溃
	//如果阻塞在子go程，那么会内存泄漏
	demoChan := make(chan int, 10)

	//var names chan string
	//
	//names <- "hello" //names是nil的，写操作会阻塞在这里

	//读
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("读取数据:", <-demoChan)
		}
	}()

	//写
	for i := 0; i < 10; i++ {
		demoChan <- i
		fmt.Println("写入数据i：", i)
	}

	time.Sleep(5 * time.Second)
}
