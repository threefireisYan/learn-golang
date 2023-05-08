package main

import "fmt"

//Person类，绑定方法：Eat,Run,Laugh,成员

type Person struct {
	name   string
	age    int
	gender string
	score  float32
}

// 在类外部绑定方法
// 如果使用指针，然后修改了this成员的值，则会修改到指针指向内存中的值
func (this *Person) Eat() {
	fmt.Println("Person is eat")
	//	类的方法，可以使用自己的成员
	this.name = "duke"
}

// 在类外部绑定方法
// 如果使用非指针，然后修改了this成员的值，则不会修改到指针指向内存中的值
func (this Person) Eat2() {
	fmt.Println("Person is eat")
	//	类的方法，可以使用自己的成员
	this.name = "duke"
}

func main() {
	lily := Person{
		name:   "lily",
		age:    20,
		gender: "女",
		score:  80,
	}

	fmt.Println("非指针修改前lily:", lily)
	lily.Eat2()
	fmt.Println("非指针修改后lily:", lily)

	fmt.Println("指针修改前lily:", lily)
	lily.Eat()
	fmt.Println("指针修改后lily:", lily)
}
