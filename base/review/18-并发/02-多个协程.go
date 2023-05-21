package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 20; i++ {
		//这里调用多个协程的时候是由cpu进行分配的，执行的时候不会按照顺序进行执行
		go hec(i)
	}

	time.Sleep(5 * time.Second)
	fmt.Println("主线程结束")

}

func hec(a int) {
	fmt.Println("hello world :", a)
}
