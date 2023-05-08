package main

import "fmt"

// 1.函数返回值在参数列表之后
// 2.如果有多个返回值，需要使用括号，多个参数之间使用，需要，分割
func test2(a int, b int, c string) (int, string, bool) {
	return a + b, c, true
}

func test3(a, b int, c string) (res int, str string, bl bool) {
	var i, j, k int
	i = 2
	j = 3
	k = 4
	fmt.Println(i, j, k)
	//	当返回值有名字的时候，可以直接写return

	res = a + b
	str = c

	return
}

// 若返回值只有一个参数，并且没有名字，那么不需要加括号
func test4() int {
	return 123
}

func main() {
	v1, s1, _ := test2(1, 2, "he")
	fmt.Println("v1:", v1)
	fmt.Println("s1:", s1)

	fmt.Println(test4())
	v2, s2, _ := test3(1, 2, "33")
	fmt.Println("v2:", v2)
	fmt.Println("s2:", s2)

}
