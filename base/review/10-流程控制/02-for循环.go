package main

import "fmt"

func main() {
	sum := 0
	//i := 0; 赋初值，i<10 循环条件 如果为真就继续执行 ；i++ 后置执行 执行后继续循环
	for i := 0; i < 10; i++ {
		sum += i
	}
	//变种
	step := 2
	for step > 0 {
		step--
		fmt.Println(step)
	}

	//	主动结束循环

}
