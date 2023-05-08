package main

import "fmt"

func main() {
	//names := [10]string{"1","2","3","4"}
	names := []string{"1", "2", "3", "4", "6"}
	for i, v := range names {
		fmt.Println("i:", i, ",v", v)
	}

	//append后names是没有变化的
	names2 := append(names, "5")
	fmt.Println("names:", names)
	fmt.Println("names2:", names2)

	fmt.Println("追加元素钱，name的长度", len(names), "，容量：", cap(names))

	//所以需要追加给自己，需要赋值给自己
	names = append(names, "5")
	fmt.Println("names:", names)
	//容量达到上限后，追加后容量会增加两倍
	fmt.Println("追加元素钱，name的长度", len(names), "，容量：", cap(names))

	names = append(names, "7")
	fmt.Println("追加元素钱，name的长度", len(names), "，容量：", cap(names))
	nums := []int{}
	for i := 0; i < 35; i++ {
		nums = append(nums, i)
		fmt.Println("len:", len(nums), ",cap:", cap(nums))
	}

}
