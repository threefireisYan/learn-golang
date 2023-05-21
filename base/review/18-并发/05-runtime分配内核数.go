package main

import (
	"fmt"
	"runtime"
	"time"
)

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
	fmt.Println("A:", time.Now())

}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
	fmt.Println("B:", time.Now())
}

func main() {
	runtime.GOMAXPROCS(2)
	go a()
	go b()
	time.Sleep(time.Second)
}
