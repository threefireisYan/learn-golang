package main

import "fmt"

func main() {
	//	浅拷贝
	//	定义一个数组
	var intTest = [5]int{1, 2, 3, 4, 5}
	//	利用数组生成一个切片
	sliceTest := intTest[2:4]
	//	赋值给另一个切片
	sliceTest2 := sliceTest
	//	更改第二个切片
	sliceTest2[0] = 99
	//	更改数值后会影响到原切片与原数组
	//	intTest: [1 2 99 4 5]
	//	sliceTest: [99 4]
	//	sliceTest2: [99 4]
	fmt.Println("intTest:", intTest)
	fmt.Println("sliceTest:", sliceTest)
	fmt.Println("sliceTest2:", sliceTest2)

	//	深拷贝
	//	定义一个数组
	var intTest2 = [5]int{1, 2, 3, 4, 5}
	//	定义一个切片（一定要使用make指明创建大小），如果用var创建一个空切片进行copy是无法将数组的值传入的 ，如下
	//var sliceTest3 []int
	sliceTest3 := make([]int, 10, 20)
	copy(sliceTest3, intTest2[1:4])
	//	更改数值后不会影响原切片和数组
	sliceTest3[0] = 99
	fmt.Println("intTest2:", intTest2)
	fmt.Println("sliceTest3:", sliceTest3)

}
