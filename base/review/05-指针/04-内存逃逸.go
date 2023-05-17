package main

// 定义一个全局变量指针
var ptr *int

func testA() {
	//定义一个x局部变量
	var x int
	x = 123
	ptr = &x
}

func testB() {
	//定义一个y局部变量指针
	var y *int
	*y = 123
}

//上述代码中，函数testA的变量x必须在堆上分配，因为他在函数退出后依然可以通过包一级的ptr变量找到，虽然他是函数内部定义的。
//在go中，这个局部变量x从函数testA中逃逸了
//相反，当testB函数返回时，变量y不再被使用，也就是可以马上被回收。因此y没有从函数testB中逃逸，编译期可以选择在栈上分配*y的存储空间，也可以在堆上分配
//然后由go语言的GC（垃圾回收机制）回收这个变量的内存空间

func main() {

}
