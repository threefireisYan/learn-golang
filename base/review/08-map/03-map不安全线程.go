package main

import "sync"

func main() {
	//var ms sync.WaitGroup
	var mutk sync.Mutex

	//	创建一个int到int的映射
	m := make(map[int]int)
	//	开启一段并发代码，创建一个goroute 协程
	go func() {
		//	不停的对map进行写入
		for true {
			mutk.Lock()
			defer mutk.Unlock()
			m[1] = 1
		}
	}()
	//	开启另一段并发代码，创建一个goroute 协程
	go func() {
		//	不停的对map进行读取
		for true {
			mutk.Lock()
			defer mutk.Unlock()
			_ = m[1]
		}
	}()

	for true {

	}

}
