package main

import "fmt"

type Sayer interface {
	Say()
}

type Eater interface {
	Eats()
}

type Dog struct {
	name string
}

func (d Dog) Say() {
	fmt.Println("汪汪汪....")
}

func (d Dog) Eats() {
	fmt.Println("吧唧吧唧...")
}

func main() {
	var sa Sayer
	var ea Eater
	var dahuang = Dog{name: "大黄"}
	sa = dahuang
	ea = dahuang
	sa.Say()
	ea.Eats()
}
