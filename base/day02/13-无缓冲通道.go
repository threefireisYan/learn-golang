package main

import (
	"fmt"
	"time"
)

func main() {
	//sync.RWMutex{}
	//设计多线程时，go语言支持使用管道、通道来解决资源同步，避免资源竞争问题
	//	试用通道，不需要进行加解锁
	//	A向通道写数据	B从管道里面读数据，go自动帮我们做好数据同步

	//	创建管道：创建一个装数字的管道
	// 此时是无缓冲管道
	numChan := make(chan int, 10)
	// 此时是有缓冲管道
	//numChan2 := make(chan int,10)

	//strChan := make(chan string)
	//创建两个go程，父写数据，子读数据

	go func() {
		for i := 0; i < 50; i++ {
			data := <-numChan
			fmt.Println("子go程1读取数据data:", data)
		}
	}()

	go func() {
		for i := 0; i < 20; i++ {
			numChan <- i
			fmt.Println("子go程2写入数据i:", i)
		}
	}()

	for i := 0; i < 30; i++ {
		//向管道中写数据
		numChan <- i
		fmt.Println("-->这是主go程，写入数据：", i)

	}

	time.Sleep(5 * time.Second)

}
