package main

import "fmt"

func main() {
	//	1.nil不是关键字或保留字
	//var nil = errors.New("hello")
	//fmt.Println(nil)

	//	nil没有默认类型
	//fmt.Printf("%T", nil)

	//	不同类型nil的指针是一样的
	//var arr []int
	//var num *int
	//
	//fmt.Printf("%p\n", arr)
	//fmt.Printf("%p", num)
	var m map[int]string
	var ptr *int
	var c chan int
	var sl []int
	var f func()
	var i interface{}
	fmt.Printf("%##v\n", m)
	fmt.Printf("%##v\n", ptr)
	fmt.Printf("%##v\n", c)
	fmt.Printf("%##v\n", sl)
	fmt.Printf("%##v\n", f)
	fmt.Printf("%##v\n", i)
}
