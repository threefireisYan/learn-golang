package main

import "fmt"

//接口需要定义方法以及方法的参数值、返回值类型
//需要类型、方法名、参数相同，且结构体方法全部对应接口，结构体方法才能正确实现接口

type Actions interface {
	Study() string
	Sleep() string
	//run() 如果接口中有，但实现结构体中没有定义，那么整个接口是不被实现的
}

type Student struct {
	name string
}

func (s Student) Study() string {
	return "learing..."
}

func (s Student) Sleep() string {
	return "sleeping..."
}

func (s Student) Eat() string {
	return "Eating..."
}

func main() {
	xiaoli := Student{name: "小李"}
	fmt.Println(xiaoli.Study())
	fmt.Println(xiaoli.Sleep())

	//先定义一个接口，然后用结构体实例化接口
	//var ac Actions
	//接口中只能少于结构体方法不能多于结构体方法,否则定义后赋值现有结构体报错
	//ac = &xiaoli
	at := Student{name: "王五"}
	//接口只能调用接口中已定义的方法，不能调用实现结构体中的其他接口中没有定义的方法
	//fmt.Println(ac.Sleep())
	fmt.Println(at.Study())

}
