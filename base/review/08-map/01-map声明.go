package main

import "fmt"

func main() {
	//	普通方式的map定义
	var mapTest1 map[string]int
	var mapTest2 map[string]int
	//	普通的赋值
	mapTest1 = map[string]int{"key1": 12, "key2": 13}
	//	需要注意这里因为map是引用类型，同样是浅拷贝赋值
	mapTest2 = mapTest1
	fmt.Println("mapTest1:", mapTest1)
	fmt.Println("mapTest2:", mapTest2)
	//	更改mapTest2赋值来的元素值
	mapTest2["key1"] = 44
	fmt.Println("mapTest1:", mapTest1)
	fmt.Println("mapTest2:", mapTest2)

	//	使用make进行定义
	mapTest3 := make(map[string]int, 10)
	fmt.Println("mapTest3:", mapTest3)

	//map不支持使用copy进行拷贝，copy仅支持切片
	//copy(mapTest3, mapTest1)
	//如需要进行赋值给新map，则需要对原map进行遍历赋值，
	/*	这只复制了 map 的结构和内容，而不是底层数据。如果 map 的值是引用类型（如切片或结构体），
		则复制后的 map 中的对应值仍然是指向相同的底层数据。如果需要深层复制，可以根据具体需求编写相应的逻辑。
	*/
	originalMap := map[string]int{"a": 1, "b": 2, "c": 3}
	// 创建一个新的空 map
	copiedMap := make(map[string]int)
	// 遍历原始 map，并复制每个键值对到新的 map
	for key, value := range originalMap {
		copiedMap[key] = value
	}
	fmt.Println("copiedMap:", copiedMap)

}
