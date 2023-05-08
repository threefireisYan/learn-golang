package main

import "fmt"

func main() {
	names := [7]string{"1", "2", "3", "4", "5", "6", "7"}
	fmt.Println("names:", names)
	//	基于定长数组创建一个新的数组
	names1 := [3]string{}
	names1[0] = names[0]
	fmt.Println("names1:", names1)
	//	切片可以基于数组，创建新的数组

	names2 := names[0:3] //左闭右开
	fmt.Println("names2:", names2)

	//	切片创建的新数组是引用内存地址的，新数组值修改，旧数组也会修改
	names2[0] = "test"
	fmt.Println("修改后的names2:", names2)
	fmt.Println("names:", names)

	//	截取最后一位，直接:右边为空
	names4 := names[0:]
	fmt.Println("names4:", names4)

	names4 = names[:]
	fmt.Println("names4:", names4)

	//	也可以切片字符串
	sub1 := "helloworld"[2:5]
	fmt.Println("sub1:", sub1)

	//创建空切片，明确切片容量
	str2 := make([]string, 10, 20)
	fmt.Println("str2:", str2, ",len:", len(str2), ",cap:", cap(str2))

	//	如果想让切片完全独立于原始数据，可以使用copy()函数来完成
	//
	//	names3 := copy(names,0:3)
	//
	//	fmt.Println(names3)
	//
	namesCopy := make([]string, len(names))
	copy(namesCopy, names[:]) //深拷贝
	fmt.Println("修改前namesCopy:", namesCopy)
	namesCopy[0] = "213"
	fmt.Println("修改后namesCopy:", namesCopy)
	fmt.Println("原names：", names)

}
