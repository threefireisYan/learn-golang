package main

import "fmt"

func main() {
	//	1.定义一个字典
	//	学生id==>学生名字 idNames
	var idNames map[int]string // 定义一个map,此时这个Map是不能直接赋值的，是空值是nil

	//  2.分配空间,使用Make,建议直接分配好固定长度空间
	idNames = make(map[int]string, 10)

	idNames[0] = "duke"

	//3.定义时直接分配空间
	idNames2 := make(map[int]string, 10)

	idNames[0] = "duke"

	for key, value := range idNames {
		fmt.Println("key:", key, ",value:", value)
	}

	for key, value := range idNames2 {
		fmt.Println("key:", key, ",value:", value)
	}

	//	5.如何确定一个Key是否存在Map中
	//	在map中不存在访问月结的问题，map认为所有的key都是有效的，所以访问一个不存在的key不会崩回，返回这个类型的零值
	//	对于int是0
	//	对于string是 空
	//	对于bool 是 false

	name9 := idNames[11]
	fmt.Println("name9:", name9)

	//无法通过value来判断一个Key是否存在
	//需要一个能够校验key是否存在的方式
	value, ok := idNames[0]
	if ok {
		//如果id=1是存在的，那么value就是Key=1对应值，ok返回true,繁殖返回零值
		fmt.Println("id=1这个key是存在的，value:", value)
	}

	value, ok = idNames[1]

	if ok {
		//如果id=1是存在的，那么value就是Key=1对应值，ok返回true,繁殖返回零值
		fmt.Println("id=2这个key是不存在的，value:", value)
	}
	//删除map中的元素
	//使用delete
	fmt.Println("删除前idNames:", idNames)
	delete(idNames, 0)
	fmt.Println("删除后idNames:", idNames)
	delete(idNames, 2)
	fmt.Println("删除不存在idNames:", idNames)

	//	并发任务处理，需要对map进行上锁

}
