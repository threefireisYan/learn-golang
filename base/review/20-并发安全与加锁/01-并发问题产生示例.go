package main

import (
	"fmt"
	"sync"
)

var x int64
//同步等待
var wg sync.WaitGroup

func add() {
	for i := 0; i < 5000; i++ {
		x = x + 1
	}
	//等待值-1
	wg.Done()
}
func main() {
	//增加等待值
	wg.Add(2)
	go add()
	go add()
	//等待值不为0时，等待
	wg.Wait()
	//等待值为0时结束等待
	fmt.Println(x)
}
