package main

func main() {
	//	创建一个int到int的映射
	m := make(map[int]int)
	//	开启一段并发代码，创建一个goroute 协程
	go func() {
		//	不停的对map进行写入
		for true {
			m[1] = 1
		}
	}()
	//	开启另一段并发代码，创建一个goroute 协程
	go func() {
		//	不停的对map进行读取
		for true {
			_ = m[1]
		}
	}()

	for true {

	}

}
