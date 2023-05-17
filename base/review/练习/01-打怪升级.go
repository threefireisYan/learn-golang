package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//	需求描述：（简化版，不需要结构体等）
//		1.运行程序能够获取输入行
//		2.进入无限循环，能够不断输入字符，获取字符
//		3.可以创建角色，输入姓名，经验值，打怪获取经验值，升级
//		4.退出游戏

//	定义全局变量：姓名，经验值，等级

var (
	name  string     //角色名称
	ex    int64  = 0 //角色经验值
	level int64  = 0 //角色等级
)

func main() {
	fmt.Println("欢迎进入简易打怪游戏，请输入您的角色名称")
	//	获取输入字符流，补货标准输入，并转换字符串
	reader := bufio.NewReader(os.Stdin)
	//	reader.ReadString仅能使用一次，可以等效于fmt.Scanf("%s",&name)
	readString, err := reader.ReadString('\n')
	if err != nil {
		//异常退出
		panic(err)
	}
	//删除最后的\n
	name = readString[:len(readString)-1]

	fmt.Printf("创建成功！\n")
	fmt.Printf("角色名称：%s \n 当前角色等级:%d \n 当前角色经验值:%d \n", name, level, ex)

	for true {
		fmt.Printf("您遇到了一只怪物！请进行您的选择：\n")
		fmt.Printf("1.战斗 \n")
		fmt.Printf("2.逃跑 \n")
		fmt.Printf("3.退出游戏 \n")
		selects, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		//去除\n以及两边空格
		chooes := strings.TrimSpace(selects)
		switch chooes {
		case "1":
			fmt.Println("您战胜了怪物！获取到了10点经验值！")
			ex += 10
			fmt.Println("当前等级:", getLevel(ex))
			fmt.Println("当前经验值:", ex)
		case "2":
			fmt.Println("您成功逃跑！")
		case "3":
			fmt.Println("已退出游戏！")
			os.Exit(1)
		default:
			fmt.Println("没有此选项再试一下吧！")
		}
	}
}

func getLevel(exs int64) int64 {
	if exs <= 10 {
		level = 1
	} else if exs > 10 && exs <= 20 {
		level = 2
	} else if exs > 20 && exs <= 30 {
		level = 3
	}
	return level
}
