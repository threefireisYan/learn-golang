package main

import "fmt"

type bag struct {
	items []int
}

// 变更结构体内部变量的方法最好是使用指针接收器
func (this *bag) svae(itemsId int) {
	this.items = append(this.items, itemsId)
	fmt.Println(this)
}

func (this bag) svaelist(itemsId []int) {
	for _, value := range itemsId {
		this.items = append(this.items, value)
	}
	fmt.Println(this)
}

func main() {
	bag1 := bag{
		items: nil,
	}

	//ints := make([]int, 10, 20)
	ints := [5]int{1, 2, 3, 4, 5}
	fmt.Println(ints)
	//ints = append(ints, 99)
	//ints = append(ints, 100)
	fmt.Println(ints)

	bag1.svae(123)
	//list没有使用指针接收器那么无法改动到bag1的变量值
	bag1.svaelist(ints[0:])

	fmt.Println(bag1.items)

}
