package main

import "fmt"

func main() {
	var whatever = [5]int{1, 2, 3, 4, 5}
	for i, _ := range whatever {
		i := i
		defer func() { fmt.Println(i) }()
	}
}
