package main

import "log"

func main() {
	var s = make([]int, 3, 5)
	log.Printf("开始：s的变量地址:%p, s的底层数组首地址:%p\n", &s, s)
	//将切片作为入参的时候不是纯引用类型入参，还是需要定义为指针类型，才可影响原始切片
	A(&s)
	log.Printf("调用A后：s的变量地址:%p, s的底层数组首地址:%p\n", &s, s)
	log.Println("原始slice s:", s)
}

func A(slice *[]int) {
	log.Printf("---A开始：参数slice的变量地址:%p, slice的底层数组首地址:%p\n", &slice, slice)
	*slice = append(*slice, 9)
	log.Println("A函数内的slice结果:", slice)
	log.Printf("---A结束：参数slice的变量地址:%p, slice的底层数组首地址:%p\n", &slice, slice)
}
