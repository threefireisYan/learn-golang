package main

import "fmt"

func test2() {
	defer func() { //第三步
		// defer panic 会打印
		fmt.Println(recover())
	}()

	defer func() { //	第二步
		panic("defer panic")
	}()
	//	第一步
	panic("test panic")
}

func main() {
	test2() //仅打印最后一个报错
}
