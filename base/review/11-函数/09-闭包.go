package main

import (
	"fmt"
)

// 创建一个玩家生成器, 输入名称, 输出生成器
func playerGen(name string) func() (string, int) {
	// 血量一直为150
	hp := 150
	// 返回创建的闭包
	return func() (string, int) {
		// 将变量引用到闭包中
		hp++
		return name, hp
	}
}

// 创建一个玩家生成器, 输入名称, 输出生成器
func playerGen1() func(string) (string, int) {
	// 血量一直为150
	hp := 150
	// 返回创建的闭包
	return func(name string) (string, int) {
		// 将变量引用到闭包中
		hp++
		return name, hp
	}
}
func main() {
	// 创建一个玩家生成器
	generator := playerGen("码神")
	// 返回玩家的名字和血量
	name, hp := generator()
	name1, hp1 := generator()
	name2, hp2 := generator()
	// 打印值
	fmt.Println(name, hp)
	fmt.Println(name1, hp1)
	fmt.Println(name2, hp2)

	generator1 := playerGen1()
	name3, hp3 := generator1("码神")
	name4, hp4 := generator1("码神1")
	name5, hp5 := generator1("码神2")
	// 打印值
	fmt.Println(name3, hp3)
	fmt.Println(name4, hp4)
	fmt.Println(name5, hp5)
	fmt.Println(name5, hp5)
}
