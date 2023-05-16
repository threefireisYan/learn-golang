package main

import (
	"fmt"
	"math"
)

func main() {
	//	go不支持隐式转换类型，只能显式转换
	var i int64
	var str string
	i = 123
	//这种赋值方式会直接报错
	//str = i
	//需要手动显式转换,但是数值转字符串会以对应的ASCII码或者Unicode码进行展示
	str = string(i)
	fmt.Println(str)
	//类型转换只能在定义正确的情况下转换成功，例如：一个取值范围较小的类型转换到一个取值范围较大的类型（int16转为int32）。
	//1.当从一个取值范围较大的类型转换到取值范围较小的类型时（将int32转换为int 16或float32转换为int），会发生精度丢失的情况
	//2.只有相同底层类型的变量之间可以进行相互转换（如将int16转换成int32类型），不同底层类型的变量相互转换时会发生编译错误（如bool类型转换为int类型）
	//如下就会报错
	//str = "111"
	//i = int64(str)

	//	输出各数值范围
	fmt.Println("int8 range:", math.MinInt8, math.MaxInt8)
	fmt.Println("int8 range:", math.MinInt16, math.MaxInt16)
	fmt.Println("int8 range:", math.MinInt32, math.MaxInt32)
	fmt.Println("int8 range:", math.MinInt64, math.MaxInt64)
	//初始化一个32位整型值
	var a int32 = 1047483647
	//输出变量的十六进制形式和十进制值
	fmt.Printf("int32:0x%x %d\n", a, a)
	//将a变量数值转换为十六位，发生数值截断(精度缺失)
	b := int16(a)
	fmt.Printf("int16:0x%x %d\n", b, b)
	//将常量保存为float32类型
	var c float32 = math.Pi
	fmt.Println(int(c))

}
