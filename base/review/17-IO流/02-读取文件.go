package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//	打开文件,返回的是文件指针
	file, err := os.Open("./17-IO流/02-读取文件.go")
	if err != nil {
		log.Printf("读取文件失败:%v", err)
		return
	}
	defer file.Close()
	//	定义接受文件读取的字节数组
	var buf [128]byte
	var content []byte
	for {
		n, err := file.Read(buf[:])
		if err == io.EOF {
			// 文件末尾
			break
		}
		if err != nil {
			log.Printf("读取文件失败:%v", err)
		}
		content = append(content, buf[:n]...)
	}
	fmt.Println(string(content))
}
