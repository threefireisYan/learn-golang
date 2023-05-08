package main

import (
	"fmt"
	"time"
)

// worker函数用于模拟工作协程，从通道jobs中读取任务并执行，如果接收到了stop通道的关闭信号，工作协程就停止工作。
func worker(id int, jobs <-chan int, stop <-chan struct{}) {
	for {
		select {
		// 从jobs通道中读取任务
		case job := <-jobs:
			fmt.Printf("worker %d start job %d\n", id, job)
			time.Sleep(time.Second) // 模拟任务耗时
			fmt.Printf("worker %d finish job %d\n", id, job)
		// 接收到了stop通道的关闭信号
		case <-stop:
			fmt.Printf("worker %d stopped\n", id)
			return
		}
	}
}

func main() {
	// 设置工作协程数和任务数
	numWorkers := 3
	numJobs := 10

	// 创建jobs通道，缓冲大小为numJobs，stop通道用于控制工作协程的停止
	jobs := make(chan int, numJobs)
	stop := make(chan struct{})

	// 创建numWorkers个工作协程，并启动它们
	for i := 1; i <= numWorkers; i++ {
		go worker(i, jobs, stop)
	}

	// 向jobs通道中写入numJobs个任务
	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}

	// 关闭stop通道，通知工作协程停止工作
	close(stop)

	// 等待工作协程全部停止，这里采用了time.Sleep的方式等待，实际使用时可以使用sync.WaitGroup等方式
	time.Sleep(5 * time.Second)
}
