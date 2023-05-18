package main

import "fmt"

func main() {
	//直接使用已定义的TESTc
	testF := TestF(fnd)
	//直接使用匿名函数
	f := TestF(func() int {
		return 123
	})

	fmt.Println(testF, f)

}

func TestF(fnd func() int) int {
	return fnd()
}

func fnd() int {
	return 332
}
