package main

func main() {
	//使用make进行分配通道空间,且定义通道具体容量
	ch1 := make(chan int, 10)
	ch1 <- 10
	//chValue := <-ch1
	//fmt.Printf("%d", <-ch1)
}
