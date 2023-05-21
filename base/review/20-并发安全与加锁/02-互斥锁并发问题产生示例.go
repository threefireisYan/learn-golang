package main

import (
	"fmt"
	"sync"
)

var x1 int64

// 同步等待
var wg2 sync.WaitGroup

// 定义一个所
var mx sync.Mutex

func add2() {
	for i := 0; i < 5000; i++ {
		mx.Lock()
		x1 = x1 + 1
		fmt.Println(x1)
		mx.Unlock()
	}
	//等待值-1
	wg2.Done()
}
func main() {
	//增加等待值
	wg2.Add(2)
	go add2()
	go add2()
	//等待值不为0时，等待
	wg2.Wait()
	//等待值为0时结束等待
	fmt.Println(x1)
}
