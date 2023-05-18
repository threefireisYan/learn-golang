package main

import (
	"fmt"
	"time"
)

func main() {
	var whatever = [5]int{1, 2, 3, 4, 5}

	for i := range whatever {
		defer fmt.Println(i)
	}

	time.Sleep(2 * time.Second)
}
