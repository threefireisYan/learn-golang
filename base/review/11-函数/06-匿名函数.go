package main

import (
	"fmt"
	"math"
)

func main() {
	//这里将一个函数当做一个变量一样的操作。
	getSqrt := func(a float64) float64 {
		return math.Sqrt(a)
	}
	fmt.Println(getSqrt(4))
	//定义时调用 带()时 funcValue赋值的是一个匿名函数返回的具体值
	funcValue := func(data int) int {
		return data
	}(100)
	fmt.Println(funcValue)
}
