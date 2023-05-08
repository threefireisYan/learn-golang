package main

import (
	"day02/08-类权限/src"
	"fmt"
)

func main() {
	s1 := src.Studnet{
		Hum: src.Human{
			Name:   "lily",
			Age:    18,
			Gender: "女",
		},
		School: "测试学习",
	}

	fmt.Println("s1 name:", s1.Hum.Name)
	fmt.Println("s1 school:", s1.School)

	//继承的时候，虽然没有定义字段名称，单会自动创建一个默认的同名字段
	//这是为了在子类中依然可以操作父类，原因是子类父类可能出现同名的字段
	t1 := src.Teacher{}
	t1.Name = "王老师"
	t1.Age = 33
	t1.Gender = "男"
	t1.Subject = "语文"
	t1.Score = 80
	fmt.Println("t1:", t1)
	t1.Eat()
}
