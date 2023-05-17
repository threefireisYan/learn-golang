package main

import (
	"fmt"
	"math"
	"strconv"
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

	//字符串与其他数据类型的转换
	//str 转 int
	newStr := "1"
	//使用strconv函数进行转义
	intValue, _ := strconv.Atoi(newStr)
	fmt.Printf("str类型转换int：%T,%d \n", intValue, intValue)

	//int 转 str
	intValue2 := 1
	newStr2 := strconv.Itoa(intValue2)
	fmt.Printf("int类型转换str：%T,%s\n", newStr2, newStr2)

	//string 转 float
	str1 := "3.14233"
	floatValue, _ := strconv.ParseFloat(str1, 64)
	fmt.Printf("str转float64:%T,%f", floatValue, floatValue)

	//float 转 string
	floatValue2 := 3.12314
	floatStr := strconv.FormatFloat(floatValue2, 'f', 2, 64)
	//4个参数，	1：要转换的浮点数
	//			2：格式标记（b、e、E、f、g、G）
	//			3：精度
	//			4：指定浮点类型（32：float32、64：float64）
	//	格式标记：
	//		'b':(-dddbp±ddd,二进制指数)
	//		'e':(-d.dddbe±ddd,十进制指数)
	//		'E':(-d.dddbE±ddd,十进制指数)
	//		'f':(-dddbp±ddd,没有指数)
	//		'g':('e':大指数,'f':其他情况)
	//		'G':('E':大指数,'f':其他情况)
	//
	//	如果格式标记为 'e','E'和'f',则prec表示小数点后的数字位数
	//	如果格式标记为 'g','G'，则prec表示总的数字尾数（整数部分+小数部门）
	fmt.Printf("float64转str:%T,%s", floatStr, floatStr)

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
