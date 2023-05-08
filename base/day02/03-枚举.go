package main

import "fmt"

// 在go语言中没有枚举类型，但是可以使用const+iota（常量累加器）来进行模拟
// 1.iota是常量组计数器
// 2.iota 从0开始，换行递增1
// 3.常量组有个特点，如果不复制，默认与上一行表达式相同
// 4.如果同一行出现两个iota，那么两个iota的值是相同的
// 5.每个常量组的iota是独立的，如果遇到const iota会重新从0开始，但可以对iota进行加减

const (
	MONDY     = iota
	TUESDAY   = iota
	WEDNESDAY = iota
	THURSDAY
	FRIDY
	SATUDAY
	SUNDAY
	M, n = iota, iota
)
const (
	MonS = iota + 1
	Ner
)

func main() {

	//var number int
	//var name string
	//var flag bool
	//	可以使用变量组来统一定义变量
	//var(
	//	number int
	//	name string
	//	flag bool
	//)

	fmt.Println(MONDY)
	fmt.Println(TUESDAY)
	fmt.Println(THURSDAY)
	fmt.Println(FRIDY)
	fmt.Println(SATUDAY)

	fmt.Println(MonS)
	fmt.Println(Ner)
}
