package main

import "fmt"

func main() {
	//普通情况下进行变量交换
	var (
		a int
		b int
		c int
	)
	a = 100
	b = 200

	c = a
	a = b
	b = c
	fmt.Println("a", a, "b", b)

	//	异或交换
	a = 100
	b = 200

	a = a ^ b
	b = b ^ a
	a = a ^ b

	fmt.Println(a, b)

	//	go语言的直接交换
	a, b = b, a

	fmt.Println(a, b)

}
