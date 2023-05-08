package main

import "fmt"

func main() {
	citys := testptr()
	fmt.Println("citys:", *citys)
}

func testptr() *string {
	city := "深圳"
	ptr1 := &city
	return ptr1
}
