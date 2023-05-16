package main

import "fmt"

func main() {
	//定义变量
	var cat int = 1
	var str string = "GO语言"
	//获取变量的内存地址
	fmt.Printf("%p,%p\n", &cat, &str)
	//定义一个指针指向变量
	//其中cat代表被取地址的变量，变量v的地址使用变量ptr进行接受，ptr的类型为*T,*T的指针类型，*代表指针
	ptr := &cat
	fmt.Printf("%p\n", ptr)
	fmt.Println(*ptr)

	//	定义指针2
	var ptr3 *string
	*ptr3 = "测试测试"
	//定义指针3
	var ptr2 = new(string)
	*ptr2 = "测试"

}
