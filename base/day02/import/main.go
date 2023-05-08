package main

import (
	_ "day02/import/add" //仅调用init函数
	"day02/import/sub"
	"fmt"
)

func main() {

	res := sub.Sub(1, 2)
	fmt.Println("res:", res)

	//如果一个包中的函数想要对外提供访问权限，
	//那么一定要首写字母大写相当于pubilc，
	//如果首字母小写则是私有函数相当于protect,只有同层级包使用，即package需要相同
	//res = add.Add(2, 4)

	fmt.Println("res:", res)
}
