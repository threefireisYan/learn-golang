package main

import "fmt"

// 将NewInt定义为Int类型
type NewInt int

// 将int取一个别名叫IntAlias
type IntAlias = int

func main() {
	//	声明a变量为NewInt类型
	var a NewInt
	//	查看a的类型名 main.NewInt
	fmt.Printf("a type %T\n", a)
	//	声明a2比阿娘为IntAlias类型
	var a2 IntAlias
	//	查看a2的类型名 int
	//IntAlias类型知会在代码中存在，编译完成时，不会有IntAlias类型
	fmt.Printf("a2 type %T\n", a2)
}
