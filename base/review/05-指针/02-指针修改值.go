package main

import "fmt"

func main() {
	//1.定义一个变量
	var test int64 = 1000
	//2.定义一个指针，指向变量test
	ptr := &test
	//3.使用*操作符，更改指针变量的内存地址中存储的值
	*ptr = 200
	//4.打印原变量
	fmt.Println("test:", test)

	//
}
