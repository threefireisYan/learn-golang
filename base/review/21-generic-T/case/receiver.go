package _case

import "fmt"

// 泛型结构体
type Mystruct[T interface{ *int | *string }] struct {
	Name string
	Data T
}

// 泛型receiver
func (myStruct Mystruct[T]) Getdata() T {
	//var i interface{} = 10
	//a,ok := i.(int)
	// 不支持断言
	//b,ok := t.(int)
	return myStruct.Data
}

func ReceiverCase() {
	data := 18
	m := Mystruct[*int]{
		Name: "sanhuo",
		Data: &data,
	}
	data1 := m.Getdata()
	fmt.Println(*data1)

	str := "asdawd"
	m2 := Mystruct[*string]{
		Name: "sanhuo",
		Data: &str,
	}
	str2 := m2.Getdata()
	fmt.Println(*str2)
}
