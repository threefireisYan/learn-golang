package main

import "fmt"

type student struct {
	name string
	age  int
}

func main() {

	zhangsan := student{
		name: "张三",
		age:  17,
	}

	fmt.Println("名字：", zhangsan)
	fmt.Println("名字：", zhangsan.name)
	fmt.Println("年龄：", zhangsan.age)

}
