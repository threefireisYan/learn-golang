package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {

	var p Point
	//p.X = 1
	//p.Y = 2
	//如果不赋值 结构体中的变量会使用零值初始化
	fmt.Printf("%v,x=%d,y=%d", p, p.X, p.Y)
}
