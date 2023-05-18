package main

import "fmt"

func main() {
	//if与else的使用
	x := 1
	if x == 0 {
		if x > 0 {
			x = 0
			fmt.Println("x:", x)
		}
		fmt.Println("x:", x)
	} else if x > 0 {
		x = 1
		fmt.Println("x:", x)
	}
	x += 1
	fmt.Println("x:", x)

	//	go语言的特性
	//	特性：x的作用域在if中，不再main的作用域中
	if x := 33; x > 0 {
		fmt.Println("x:", x)
	}

}
