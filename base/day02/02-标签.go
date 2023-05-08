package main

import "fmt"

func main() {
	//	标签
	//	goto ==>下次进入循环时，i不会保存之前的状态，会重新计算
	//	break ==>直接跳出指定位置
	//	continue ==>下次进入循环时，i会保存之前的状态
	//	标签的名字是自定义的，后面加上冒号
LABEL1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				//goto LABEL1
				//continue LABEL1
				break LABEL1
			}
			fmt.Println("i:", i, ",j:", j)
		}
	}

	fmt.Println("end")

}
