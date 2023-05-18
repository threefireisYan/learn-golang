package main

import "fmt"

func main() {
	//	调用函数
	TestA()
	TestB(32)
	c := TestC()
	fmt.Println(c)
	d := LibTestD(32, 112)
	fmt.Println(d)
	a, e := TestE(32, 112)
	fmt.Println(a, e)
}

// 不带参数的函数
func TestA() {
	fmt.Println("TestA")
}

// 带实际参数的函数
func TestB(a int) {
	fmt.Println("TestB,a:", a)
}

// 带返回值的参数的函数
func TestC() int {
	fmt.Println("TestC")
	return 99
}

// 带多实际参数的函数
func LibTestD(a int, b int) int {
	fmt.Println("TestC")
	return a + b
}

// 带多返回参数的函数
func TestE(a int, b int) (int, string) {
	fmt.Println("TestC")
	return a + b, "test"
}
