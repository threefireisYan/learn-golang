package main

import "fmt"

func main() {
	a, b := 100, 200
	c := sum(a, b)
	fmt.Println(c)
}

func sum(a int, b int) int {
	return a + b
}
