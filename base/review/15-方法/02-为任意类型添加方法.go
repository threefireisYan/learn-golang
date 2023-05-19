package main

import "fmt"

// 定义一个自定义类型（可以不止是int，可以为string、map等）
type Myint int64

// 为自定义类型添加方法
func (m Myint) getIntsZero() int {
	return 0
}

func main() {
	var a Myint
	zero := a.getIntsZero()
	fmt.Println(zero)
}
