package main

import (
	"fmt"
	"sync"
)

func main() {
	//	sync.Map不能使用make创建
	var scene sync.Map
	//	将键值对保存到sync.Map
	scene.Store("name", 11)
	scene.Store("age", "sda")
	//	从sync.Map中取值
	fmt.Println(scene.Load("name"))
	//	删除sync.Map中的值
	//scene.Delete("name")
	//	遍历所有sync.Map的键值对
	//	遍历需要提供一个匿名函数，参数值为k,y,类型为interface{},每次range()在遍历一个元素时，都会调用和这个匿名函数把结果返回
	scene.Range(func(key, value any) bool {
		fmt.Println("iterate:", key, value)
		return true
	})

	//	开启一段并发代码，创建一个goroute 协程
	go func() {
		//	不停的对map进行写入
		for true {
			scene.Store("name", 11)
		}
	}()
	//	开启另一段并发代码，创建一个goroute 协程
	go func() {
		//	不停的对map进行读取
		for true {
			scene.Load("name")
		}
	}()

	for true {

	}

}
