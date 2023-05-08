package main

import (
	"fmt"
	"runtime"
	"time"
)

// GOEXIT 提前退出当前线程
// return 返回当前函数
// exit 退出当前进程
func main() {
	go func() {
		go func() {
			fmt.Println("这是子线程内部的子线程！")
			//return
			//os.Exit(-1)
			runtime.Goexit() //退出当前子线程
			fmt.Println("还打印吗？")
		}()
		fmt.Println("子线程结束")
	}()
	fmt.Println("这是主线程")

	time.Sleep(5 * time.Second)

}
