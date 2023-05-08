package src

import "fmt"

//在Go语言中，权限都是通过首写字母的大小来控制
//1.import ==> 如果包名不同，只有首字母大写才是
//2.对于类里面的成员、方法 ==> 只有首字母大写才能在其他包使用

type Human struct {
	Name   string
	Age    int
	Gender string
	Score  float32
}

func (this *Human) Eat() {
	fmt.Println("this is :", this.Name)
}

type Studnet struct {
	Hum    Human //包含Human类型的变量
	School string
}

type Teacher struct {
	Human   //直接写Human类型没有字段名称，代表继承
	Subject string
}
