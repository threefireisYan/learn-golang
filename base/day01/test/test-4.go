package main

import "fmt"

func main() {
	name := "duke"
	usage := `./a.out<option>
				-h help
				xxx`
	fmt.Println("name:", name)
	fmt.Println("usage:", usage)

	l1 := len(name)

	fmt.Println("l1:", l1)

	for i := 0; i < len(name); i++ {
		fmt.Println(i, name[i])
	}
	str1 := "hell"
	str2 := "world"
	str3 := str1 + str2
	fmt.Println(str3)

	const address = "nanjing"
	//address = "beijing"
	fmt.Println("address", address)

}
