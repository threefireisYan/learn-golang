package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func wr() {
	// 参数2：打开模式，所有模式d都在上面
	// 参数3是权限控制
	// w写 r读 x执行   w  2   r  4   x  1
	//特殊权限位，拥有者位，同组用户位，其余用户位
	file, err := os.OpenFile("./17-IO流/test.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return
	}
	defer file.Close()
	// 获取writer对象
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("hello\n")
	}
	// 刷新缓冲区，强制写出
	writer.Flush()
}

func re() {
	file, err := os.Open("./17-IO流/test.txt")
	if err != nil {
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		//按行读取
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}
		fmt.Println(string(line))
	}

}

func main() {
	re()
}
