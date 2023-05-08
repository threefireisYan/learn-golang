package main

import "fmt"

func main() {
	channlOne := make(chan int, 10)

	go func() {
		defer close(channlOne)
		for i := 0; i < 10; i++ {
			channlOne <- i
		}
	}()

	for {
		v, ok := <-channlOne
		if !ok {
			fmt.Println("通道已经关闭")
			break
		}
		fmt.Println("v:", v)
	}
	fmt.Println("END")
	fmt.Println("END2")

}
