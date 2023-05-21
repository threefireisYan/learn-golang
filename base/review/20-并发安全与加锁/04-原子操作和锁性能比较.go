package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var x4 int64
var l sync.Mutex
var wg33 sync.WaitGroup

// 原子操作仅限于对基本数据类型操作
// 普通版加函数
func add3() {
	// x = x + 1
	x4++ // 等价于上面的操作
	wg33.Done()
}

// 互斥锁版加函数
func mutexAdd() {
	l.Lock()
	x4++
	l.Unlock()
	wg33.Done()
}

// 原子操作版加函数
func atomicAdd() {
	atomic.AddInt64(&x4, 1)
	wg33.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 10000; i++ {
		wg33.Add(1)
		// go add()       // 普通版add函数 不是并发安全的
		// go mutexAdd()  // 加锁版add函数 是并发安全的，但是加锁性能开销大
		go atomicAdd() // 原子操作版add函数 是并发安全，性能优于加锁版
	}
	wg33.Wait()
	end := time.Now()
	fmt.Println(x4)
	fmt.Println(end.Sub(start))
}
