package main

import "fmt"

// 1.实现go多态，需要实现定义接口
// 例子：
type IAttack interface {
	//接口的函数可以有多个，但是只能有函数原型，不可以有实现
	Attack()
}

type HumanLowLevel struct {
	name  string
	level int
}

func (this *HumanLowLevel) Attack() {
	fmt.Println("名称：", this.name, ",等级：", this.level)
}

type HumanHightLevel struct {
	name  string
	level int
}

func (this *HumanHightLevel) Attack() {
	fmt.Println("名称：", this.name, ",等级：", this.level)
}

func main() {

	//定义一个包含Attack的接口变量
	var player IAttack

	lowLevelHuman := HumanLowLevel{
		name:  "大卫",
		level: 1,
	}

	hightLeveHuman := HumanHightLevel{
		name:  "米沙",
		level: 10,
	}

	lowLevelHuman.Attack()
	hightLeveHuman.Attack()
	//对player赋值，接口需要使用指针类型来赋值(因为struct是用的指针类型)
	player = &lowLevelHuman
	player.Attack()

	player = &hightLeveHuman
	player.Attack()

}
