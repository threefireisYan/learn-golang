package main

import (
	"fmt"
	"sync"
)

var x int64

// 同步等待
var wg sync.WaitGroup

func add() {
	for i := 0; i < 50; i++ {
		x = x + 1
		fmt.Println("此时x:", x)
	}
	//等待值-1
	wg.Done()
}
func main() {
	//增加等待值
	wg.Add(2)
	//创建两个协程，同时对全局变量x进行操作
	//有可能出现以下并发问题，协程1抢占到cpu的时间片在进行循环第30次把 x赋值为 200，这时时间片结束，
	//协程2获取到时间片，开始进行第一次循环，此时将x赋值为了201，并立刻被协程1抢回时间片。
	//协程1抢回时间片，进行第31次循环，此时x已经被协程2变更为了201，是不是x在协程1的基础上又被多增加了1？
	go add()
	go add()
	//等待值不为0时，等待
	wg.Wait()
	//等待值为0时结束等待
	fmt.Println(x)
}
