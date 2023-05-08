package main

import "fmt"

func main() {
	var name string
	name = "123"
	fmt.Println(name)

	name_test := "4124"
	fmt.Println(name_test)
	test1(12, "string")
	i, j := 10, 21
	i, j = j, i
	fmt.Println(i, j)
}

func test1(a int, b string) {
	fmt.Println(a, b)
}
