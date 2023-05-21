package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x2     int64
	wg23   sync.WaitGroup
	lock   sync.Mutex
	rwlock sync.RWMutex
)

func write() {
	// lock.Lock()   // 加互斥锁
	rwlock.Lock() // 加写锁
	x2 = x2 + 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	rwlock.Unlock()                   // 解写锁
	// lock.Unlock()                     // 解互斥锁
	wg23.Done()
}

func read() {
	// lock.Lock()                  // 加互斥锁
	rwlock.RLock()               // 加读锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	rwlock.RUnlock()             // 解读锁
	// lock.Unlock()                // 解互斥锁
	wg23.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg23.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		wg23.Add(1)
		go read()
	}

	wg23.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}
