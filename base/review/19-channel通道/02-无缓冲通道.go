package main

import (
	"fmt"
)

func re(rec chan int) {
	x := <-rec
	fmt.Println(x)
}

func main() {
	//使用make进行分配通道空间,但不定义通道具体容量

	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
		//close(c)  关闭一个已经关闭的通道会引发panic
	}()
	for {
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}

	//需要使用协程对通道进行读取，才会不发生死锁
	//协程最好在通道写入数值之前进行
	ch1 := make(chan int)
	go re(ch1)
	ch1 <- 10
	//chValue := <-ch1

	fmt.Println("end")
	//time.Sleep(2 * time.Second)
}
