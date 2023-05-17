package main

import "fmt"

func main() {
	//	定义一个map
	mapTest1 := make(map[string]string)
	mapTest1["name"] = "张三"
	mapTest1["age"] = "14"
	fmt.Println(mapTest1)
	//	map元素的删除
	delete(mapTest1, "name")
	fmt.Println(mapTest1)

}
