package main

import "fmt"

func main() {
	names, m, num := f1()
	fmt.Println(names, m, num)
}

func f1() (names []string, m map[string]int, num int) {
	m = make(map[string]int)
	m["k1"] = 2
	return
}
