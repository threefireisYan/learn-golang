package main

import (
	"context"
	"os"
	"os/signal"
	_case "review/21-generic-T/case"
)

/*
	1.什么是泛型
		泛型即开发过程中编写适用于所有类型的模板，只有在具体使用的时候才能确定其真正的类型
	2.泛型的作用于应用场景
		1>增加代码的复用，从同类型的复用到不同类型的代码复用
		2>应用于不同类型间代码复用的场景，即不同类型需要写相同的处理逻辑时，最适合用泛型
	3.泛型的利弊
		1>提高了代码复用率，提高了变成效率
		2>不同类型间代码复用性，使得代码风格更加优雅
		3>增加了编译器的负担，降低了编译效率
	4.golang的泛型的使用
		1>泛型函数
		2>泛型类型
		3>泛型接口
		4>泛型结构体
		5>泛型receiver
	5.泛型的限制
		1>匿名结构体与匿名函数不支持泛型
		2>不支持类型断言
		3>不支持泛型方法，只能通过receiver来实现方法的泛型处理
		4>~后的类型必须为基本类型，不能为接口类型

*/

func main() {
	_case.SimpleCase()
	_case.CusNumTCase()
	_case.BuiltInCase()
	_case.TTypeCase()
	// 使用泛型切片、通道、类型
	_case.TTypeCase1()
	_case.InterfaceCase()
	_case.ReceiverCase()
	//避免主携程过早退出
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()
	<-ctx.Done()
}
