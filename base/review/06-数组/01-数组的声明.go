package main

import "fmt"

func main() {
	//	定义一个数组，数组的长度是固定的
	var arr [10]int
	var arr1 = [10]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	arr3 := []int{1, 2, 3}
	//	初始化定义下标为2的位置值为3
	arr4 := []int{2: 3}
	fmt.Println(arr[0])
	fmt.Println(arr1[0])
	fmt.Println(arr2[0])
	fmt.Println(arr3[0])
	fmt.Println(arr4[0])
	fmt.Println(arr4[2])
	//	循环遍历
	for key, value := range arr2 {
		fmt.Println("key-value:", arr2[key])
		fmt.Println("value:", value)
	}
	//	定义多维数组
	var arrList [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	var arrList2 [2][3]int = [2][3]int{1: {4, 5, 6}}

	fmt.Println("多维数组1:", arrList[0])
	fmt.Println("多维数组2:", arrList[0][2])
	fmt.Println("多维数组3:", arrList2[0])
	fmt.Println("多维数组3:", arrList2[1])
	fmt.Println("多维数组3:", arrList2[1][2])

	for key2, _ := range arrList2 {
		fmt.Println("多维数组1:", arrList2[key2])
	}

}
