package _case

import "fmt"

// 基本接口的定义,可用于变量的定义
type ToString interface {
	String() string
}

//var s ToString

func (u user) String() string {
	return fmt.Sprintf("ID:%d,Name:%s,Age:%d", u.ID, u.Name, u.Age)
}
func (addr address) String() string {
	return fmt.Sprintf("ID:%d,Province:%s,City:%s", addr.ID, addr.Province, addr.City)
}

// 泛型接口的定义
type GetKey[T comparable] interface {
	any
	Get() T
}

//不可这样定义
//var s GetKey[comparable]

func (u user) Get() int64 {
	return u.ID
}

func (addr address) Get() int {
	return addr.ID
}

// 列表转集合
func listToMap[k comparable, T GetKey[k]](list []T) map[k]T {
	mp := make(MapT[k, T], len(list))
	for _, data := range list {
		mp[data.Get()] = data
	}
	return mp
}

func InterfaceCase() {
	userList := []GetKey[int64]{
		user{
			ID:   1,
			Name: "三火",
			Age:  20,
		},
		user{
			ID:   2,
			Name: "三火2",
			Age:  21,
		},
	}
	addressList := []GetKey[int]{
		address{
			ID:       1,
			Province: "三火",
			City:     "20",
		},
		address{
			ID:       2,
			Province: "三火1",
			City:     "21",
		},
	}
	userMp := listToMap[int64, GetKey[int64]](userList)
	fmt.Println(userMp)
	addrMp := listToMap[int, GetKey[int]](addressList)
	fmt.Println(addrMp)

}
