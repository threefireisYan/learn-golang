package main

import (
	"fmt"
	"os"
)

func main() {
	// 新建文件
	file, err := os.Create("./17-IO流/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	//此写入会直接覆盖原内容
	for i := 0; i < 5; i++ {
		file.WriteString("ab\n")
		file.Write([]byte("cd\n"))
	}
}
