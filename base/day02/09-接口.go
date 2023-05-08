package main

import "fmt"

//在go语言中，关键字 interface来代表接口
//interface不仅是处理多态的，可以接受任一的数据类型，类似于void

func main() {
	var i, j, k interface{}
	names := []string{"123", "456"}
	i = names
	fmt.Println("切片数组：", i)

	age := 20
	j = age
	fmt.Println("数字：", j)

	str := "hello"
	k = str
	fmt.Println("字符串：", k)

	//	 判断一个变量是否为某个类型
	//values, flag := k.(int)
	//if !flag {
	//	fmt.Println("k不是int类型")
	//} else {
	//	fmt.Println("k是int类型", values)
	//}
	//常用场景，把interface当成一个函数的参数，（类似于print）,使用switch来判断用户输入的不同类型
	//	根据不同类型，做相应逻辑处理
	array := make([]interface{}, 3)
	array[0] = 1
	array[1] = "hello world"
	array[2] = true

	for _, value := range array {
		switch v := value.(type) {
		case int:
			fmt.Println("当前类型为Int,内容为：", value)
		case string:
			fmt.Println("当前类型为String,内容为：", value)
		case bool:
			fmt.Println("当前类型为Bool,内容为：", v)
		default:
			fmt.Println("不是数据类型")
		}
	}

}
