package sub

import "fmt"

// 1.init函数没有参数，没有返回值
// 2.一个包中有多个init时，调用顺序是不确定的
// 3.init是不允许手动显示调用的（会报错）
// 4.有时引用一个包可能只想使用包中的init函数（mysql的init对驱动进行初始化）
// 5.不想使用包中其他函数，为了防止编译器报错，可以使用_使用
// import (
//
//	_ "day02/import/add" //仅调用init函数
//	"day02/import/sub"
//	"fmt"
//
// )
func init() {
	fmt.Println("this is init() in package Sub")
}

func init() {
	fmt.Println("this is init()2 in package Sub")
}

func Sub(a, b int) int {
	fmt.Println(test4(1, 4))
	return a - b
}
