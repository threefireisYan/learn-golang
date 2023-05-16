package main

import "fmt"

func main() {
	//	[]byte和string可以通过强制类型转换
	s1 := "localhost:8080"
	//	强制类型转换 string to byte
	strByte := []byte(s1)

	//	下标修改
	strByte[len(s1)-1] = '1'
	fmt.Println(strByte)

	//	强制类型转换 []byte to string
	s2 := string(strByte)
	fmt.Println(s2)
}
