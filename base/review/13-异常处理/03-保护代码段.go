package main

import "fmt"

func test3(x, y int) {
	var z int

	func() {
		defer func() {
			if recover() != nil {
				z = 0
			}
		}()
		panic("test panic")
		z = x / y
		return
	}()

	fmt.Printf("x / y = %d\n", z)
}

func main() {
	test3(2, 1)
}
