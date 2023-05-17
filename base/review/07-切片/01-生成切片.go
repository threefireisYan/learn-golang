package main

import "fmt"

func main() {
	var a []int64 = []int64{1, 2, 3, 4, 5, 6, 6}
	b := a[2:4]
	//输出[3,4] 符合左闭右开原则
	fmt.Println(b)
	//	开始缺省
	fmt.Println(a[:4])
	//	结束缺省
	fmt.Println(a[2:])
	//	同时缺省
	fmt.Println(a[:])
	//	都为0，空切片
	fmt.Println(a[0:0])
	//	超界 会编译时报错
	//fmt.Println(a[0:9])

	//	直接声明字符串切片
	var strList []string
	//	声明整型切片
	var numList []int
	//	声明一个空切片
	var numListEmpty = []int{}
	//	输出三个切片
	fmt.Println(strList, numList, numListEmpty)
	//	输出是三个切片大小
	fmt.Println(len(strList), len(numList), len(numListEmpty))
	//	切片判定空的结果
	fmt.Println(strList == nil)
	fmt.Println(numList == nil)
	fmt.Println(numListEmpty == nil)

	//	对切片追加元素,形成的新的切片或对原切片进行追加元素
	strings := append(strList, "test")
	strList = append(strList, "test")
	strList = append(strList, "test2")
	fmt.Println("strings", strings)
	fmt.Println("strList", strList)

	//	使用make进行分配切片空间
	intTest := make([]int, 10, 10)
	//	打印切片长度，打印切片容量
	fmt.Println("长度：", len(intTest), "容量", cap(intTest))
	for i := 0; i < 22; i++ {
		//	为定义的切片追加元素
		intTest = append(intTest, i)
		//	打印其长度与容量，发现长度增加到了11/21后，容量会以当前容量*2（仅限于容量不大的情况，如果后续容量达到2GB，则可能是当前容量*1.5的形式扩容）
		fmt.Println("长度：", len(intTest), "容量", cap(intTest))
		fmt.Println("intTest:", intTest)
	}

	//	切片从数组中生成时，容量以起始位置至数组末尾。
	var number4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	myslice := number4[4:6]
	//	切片虽然时定义的数组的索引下标4~6，但实际上会自动将容量分配为切片起始下标4~len(number4)大小
	//	myslice 为 [5,6] len为2 cap为6 （10-4）
	fmt.Println("myslice:", myslice)
	fmt.Println("长度:", len(myslice), "，容量：", cap(myslice))
	//	重新定义切片后
	myslice = myslice[:cap(myslice)]
	fmt.Println("重新定义myslice:", myslice)
	//	原因是切片是属于指针的一种，是一种引用类型
	//		myslice引用的是number4的内存地址，所以重新进行自定义自身后访问myslice[3]是可以获取到值的
	fmt.Println("myslice:", myslice[3])

}
