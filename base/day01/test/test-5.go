package main

import "fmt"

func main() {
	//	定义数组
	nums := []int{12, 3, 1, 2, 4}
	fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		fmt.Println(i, nums[i])
	}

	for key, value := range nums {
		value++
		fmt.Println("key:", key, ",value:", value, ",nums:", nums[key])
	}

	for _, value := range nums {
		value++
		fmt.Println(",value:", value)
	}

}
