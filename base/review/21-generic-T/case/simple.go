package _case

import "fmt"

func SimpleCase() {
	var a, b = 3, 4
	var c, d float64 = 5, 6
	fmt.Println("不适用泛型，数字比较：", getMaxNumInt(a, b))
	fmt.Println("不适用泛型，数字比较：", getMaxNumFloat(c, d))
	//情况简单，编译器可以自行判断传入的值是什么类型
	fmt.Println("适用泛型，数字比较：", getMaxNum(a, b))
	//若情况复杂（比如自定结构体等）,需要显示指定传入的类型
	fmt.Println("适用泛型，数字比较：", getMaxNum[float64](c, d))
}

func getMaxNumInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMaxNumFloat(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// [泛型 约束]
func getMaxNum[T int | float64](a, b T) T {
	if a > b {
		return a
	}
	return b
}

type MyStruct[T interface{ *int | *float64 }] struct {
}

type CusNumT interface {
	// Get() string
	// Set(s string)
	// ~int64 表示包含 支持 uint8/int32/float64与int64的衍生类型
	// ~ 表示支持类型的衍生类型
	// | 表示取并集
	// 多行之间取交集
	// 如果多行没有任何交集，那么这个自定义泛型则不生效
	uint8 | int32 | float64 | ~int64
	int32 | float64 | ~int64 | uint64
}

// MyInt64是int64的衍生类型，是具有基础类型int64的新类型，与int64是不同的类型
type MyInt64 int64

// MyInt32 为 int32 的别名，与衍生类型不同的是 MyInt32与int32的完全相同的类型
type MyInt32 = int32

func CusNumTCase() {
	var a, b int32 = 3, 4
	var a1, b1 MyInt32 = a, b
	fmt.Println("自定义泛型，数字比较：", getMaxCusNum(a, b))
	fmt.Println("自定义泛型，数字比较：", getMaxCusNum(a1, b1))
	//情况简单，编译器可以自行判断传入的值是什么类型

	var c, d float64 = 5, 6
	fmt.Println("自定义泛型，数字比较：", getMaxCusNum(a, b))
	var e, f int64 = 7, 8
	var g, h MyInt64 = MyInt64(e), MyInt64(f)
	//若情况复杂（比如自定结构体等）,需要显示指定传入的类型
	fmt.Println("自定义泛型，数字比较：", getMaxCusNum[float64](c, d))
	fmt.Println("自定义泛型，数字比较：", getMaxCusNum[int64](e, f))
	fmt.Println("自定义泛型，自定义类型数字比较：", getMaxCusNum[MyInt64](g, h))
}

// [泛型 约束]
func getMaxCusNum[T CusNumT](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// 内置类型
func BuiltInCase() {
	var a, b string = "abc", "abc"
	fmt.Println("内置comparable泛型类型比较：", getBuiltInComparable(a, b))
	var c, d float64 = 6, 7
	fmt.Println("内置comparable泛型类型比较：", getBuiltInComparable(c, d))
	var f = 100.214
	printBuiltInAny(f)

}

func getBuiltInComparable[T comparable](a, b T) bool {
	// comparable 类型仅支持 两种种操作
	// == !=
	if a == b {
		return true
	}
	return false
}

func printBuiltInAny[T any](a T) {
	fmt.Println("内置类型any", a)
}
