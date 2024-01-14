package main

import "fmt"

type Man struct {
	Name string
}

func (man Man) Married1() {
	man.Name = "Mr. " + man.Name
}

func (man *Man) Married2() {
	man.Name = "Mr. " + man.Name
}

func main() {
	budi := Man{Name: "Budi"}
	budi.Married1()
	fmt.Println(budi)

	hari := Man{Name: "Hari"}
	hari.Married2()
	fmt.Println(hari)
}
