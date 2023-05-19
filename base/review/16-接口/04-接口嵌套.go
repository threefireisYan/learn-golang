package main

import "fmt"

type xyj interface {
	haier
	geli
}

type haier interface {
	run()
}

type geli interface {
	run()
}

type geliClasse struct {
	name string
}

type haierClasse struct {
	name string
}

func (g geliClasse) run() {
	fmt.Println("运行格力")
}

func (h haierClasse) run() {
	fmt.Println("运行海尔")
}

func main() {
	var xyjName1 xyj
	var xyjName2 xyj
	var geliInter geli
	var haierInter haier
	var gelis = geliClasse{name: "格力型号1"}
	var haiers = haierClasse{"海尔型号1"}

	geliInter = gelis
	haierInter = haiers
	xyjName1 = geliInter
	xyjName2 = haierInter

	xyjName1.run()
	xyjName2.run()

}
