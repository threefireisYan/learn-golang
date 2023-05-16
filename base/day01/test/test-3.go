package main

import "fmt"

func main() {
	//
	//	gc垃圾回收机制，开发人员不需要手动释放内存
	//	c语言不允许返回栈上的指针，go语言可以返回，程序会在编译的时候就会确认变量的分配位置
	//	编译的时候，如果发现有必要的，就讲变量分配到堆上

	name := "test"
	prt := &name
	fmt.Println("name:", *prt)
	fmt.Println("name ptr:", prt)

	//使用new关键字定义指针
	name2 := new(string)
	*name2 = "test2"
	fmt.Println(name2)
	fmt.Println(*name2)
	testptr := test()
	fmt.Println("testptr:", testptr)
	fmt.Println("city : ", *testptr)

	//	判断空指针， c:null c++:nullptr  go:nil
	if testptr == nil || *testptr == "nanjing" {
		fmt.Println("空指针")
	} else {
		fmt.Println("非空")
	}
}

func test() *string {
	city := "nanjing"
	ptr1 := &city
	return ptr1
}
