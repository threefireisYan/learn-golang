package main

import "fmt"

//1.当管道写满了，会写阻塞
//2.当缓冲器读完了，读阻塞
//3.如果管道没有使用make分配空间，管道默认是nil
//4.从nil的管道读取数据，写入数据，都会阻塞（不会崩溃）
//5.从一个已经close的管道读取数据时，会返回零值（不会崩溃）
//6.关闭一个已经close的管道，程序会崩溃
//7.关闭管道的动作，一定要在写端，不应该放在读端
//8.向已经close的管道写数据时，（会崩溃）
//9.读和写的次数一定要对等，否则
//	在多个go程中：资源泄漏
//	在主go程中：程序崩溃

func main() {

	demoChan := make(chan int, 10)

	go func() {
		for i := 0; i < 30; i++ {
			demoChan <- i
			fmt.Println("写入数据：", i)
		}
		fmt.Println("数据已经全部写入完毕，准备关闭管道")
		close(demoChan)
		//demoChan <- 10  已关闭的管道再写入数据会报错
	}()

	//for range不知道管道是否已经写完，所以会一直等待
	//在写入端，将管道关闭，for range遍历关闭的管道，会退出
	for v := range demoChan {
		fmt.Println("读取数据：", v)
	}

	fmt.Println("END")

}
