package main

import "fmt"

type User struct {
	id   int
	name string
}

type User1 struct {
	id   int
	name string
}

type Manager struct {
	User
	User1
	title string
}

//有多个继承的时候，方法最好不要有相同名称，否则报错

func (u *User) ToString() string {
	return fmt.Sprintf("User: %p, %v", u, u)
}

func (u *User1) ToString() string {
	return fmt.Sprintf("User: %p, %v", u, u)
}

//func (self *Manager) ToString() string {
//	return fmt.Sprintf("Manager: %p, %v", self, self)
//}

func main() {
	m := Manager{User{1, "Tom"}, User1{1, "Tom"}, "Administrator"}

	//有多个父类时，这里调用会报错
	//fmt.Println(m.ToString())
	//但是可以使用
	fmt.Println(m.User.ToString())
	fmt.Println(m.User1.ToString())
}
