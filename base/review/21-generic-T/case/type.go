package _case

import "fmt"

// 泛型类型
type user struct {
	ID   int64
	Name string
	Age  uint8
}

type address struct {
	ID       int
	Province string
	City     string
}

// 集合转列表
func mapTolist[k comparable, T any](mp map[k]T) []T {
	list := make([]T, len(mp))
	var i int
	for _, data := range mp {
		list[i] = data
		i++
	}
	return list
}

func myPrintln[T any](ch chan T) {
	for data := range ch {
		fmt.Println(data)
	}
}

func TTypeCase() {
	userMp := make(map[int64]user, 0)
	userMp[1] = user{
		ID:   1,
		Name: "sanhuo",
		Age:  18,
	}
	userMp[2] = user{
		ID:   2,
		Name: "sanhuo2",
		Age:  19,
	}

	userList := mapTolist[int64, user](userMp)
	ch := make(chan user)
	go myPrintln(ch)

	for _, u := range userList {
		ch <- u
	}
	addrMp := make(map[int64]address, 0)
	addrMp[1] = address{
		1,
		"sanhuo",
		"shanghai",
	}
	addrMp[2] = address{
		2,
		"sanhuo2",
		"shanghai2",
	}

	addrList := mapTolist[int64, address](addrMp)
	ch2 := make(chan address)
	go myPrintln(ch2)

	for _, u1 := range addrList {
		ch2 <- u1
	}
}

// 泛型切片的定义
type List[T any] []T

// 泛型定义集合类型
// 声明两个泛型，分别为 k、v
type MapT[k comparable, v any] map[k]v

// 泛型定义通道类型
type Chan[T any] chan T

func TTypeCase1() {
	userMp := make(MapT[int64, user], 0)
	userMp[1] = user{
		ID:   1,
		Name: "sanhuo",
		Age:  18,
	}
	userMp[2] = user{
		ID:   2,
		Name: "sanhuo2",
		Age:  19,
	}

	var userList List[user]
	userList = mapTolist[int64, user](userMp)
	ch := make(Chan[user])
	go myPrintln(ch)

	for _, u := range userList {
		ch <- u
	}

	addrMp := make(MapT[int64, address], 0)
	addrMp[1] = address{
		1,
		"sanhuo",
		"shanghai",
	}
	addrMp[2] = address{
		2,
		"sanhuo2",
		"shanghai2",
	}
	var addrList List[address]
	addrList = mapTolist[int64, address](addrMp)
	ch2 := make(chan address)
	go myPrintln(ch2)

	for _, u1 := range addrList {
		ch2 <- u1
	}
}
