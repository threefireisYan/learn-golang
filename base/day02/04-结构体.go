package main

import "fmt"

// c语言里面可以使用 typedef int MyInt
type MyInt int

type Person2 struct {
	name   string
	age    int
	gendes string
	score  float32
}

func main() {
	var i, j MyInt

	i, j = 10, 20

	fmt.Println("i+j:", i+j)

	lily := Person2{
		name:   "lily",
		age:    20,
		gendes: "男",
		score:  99,
	}
	//在定义期间，对结构体赋值，如果每个字段都赋值了，那么字段的名字可以直接省略不写
	//如果只对局部变量赋值，呢么必须明确指定变量名称
	duke := Person2{
		name: "duke",
		//age:    0,
		//gendes: "",
		//score:  0,
	}
	duke.score = 89
	duke.gendes = "男"
	fmt.Println(duke)
	fmt.Println("lily", lily.name, lily.gendes, lily.score, lily.age)
	str := &lily
	fmt.Println("指针str.name：", str.age, str.name)
	fmt.Println("指针(*str).name：", (*str).name)
}
