package main

import "fmt"

func main() {
	//使用make进行分配通道空间
	//ch1 := make(chan int, 10)
	//ch1 <- 10
	////chValue := <-ch1
	//fmt.Printf("%d", <-ch1)

	ch2 := make(chan int, 10)

	for i := 0; i < 10; i++ {
		ch2 <- i

	}
	close(ch2)
	for value := range ch2 {
		fmt.Println(value)
	}

}
