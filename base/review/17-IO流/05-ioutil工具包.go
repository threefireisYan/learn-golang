package main

import (
	"fmt"
	"io/ioutil"
)

func wr2() {
	err := ioutil.WriteFile("./17-IO流/test.txt", []byte("码神之路"), 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func re2() {
	content, err := ioutil.ReadFile("./17-IO流/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))
}

func main() {
	re2()
}
