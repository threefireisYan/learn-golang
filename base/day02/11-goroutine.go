package main

import (
	"fmt"
	"time"
)

// 这个将用于子线程
func display(nums int) {
	count := 1
	for true {
		fmt.Println("===>这是子go程:", nums, "当前值：", count)
		count++
		time.Sleep(1 * time.Second)
	}
}

func main() {
	//	启动子go程
	//go display()
	//
	//go func() {
	//	count := 1
	//	for true {
	//		fmt.Println("===>这是子go程:", count)
	//		count++
	//		time.Sleep(1 * time.Second)
	//	}
	//}()
	//	主go程

	for i := 0; i < 3; i++ {
		go display(i)
	}

	count := 1
	for true {
		fmt.Println("这是主go程:", count)
		count++
		time.Sleep(1 * time.Second)
	}

}
