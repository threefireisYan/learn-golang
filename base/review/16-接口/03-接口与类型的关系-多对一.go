package main

import "fmt"

type Cat struct {
	name string
}

func (c Cat) Say() {
	fmt.Println("汪汪汪....")
}

func (c Cat) Eats() {
	fmt.Println("吧唧吧唧...")
}

func main() {
	var sa Sayer
	var sa2 Sayer
	var dahuang = Dog{name: "大黄"}
	var daju = Cat{name: "大橘"}
	//狗和猫都可以调用同一个接口，并实例化
	sa = dahuang
	sa2 = daju
	sa.Say()
	sa2.Say()
}
