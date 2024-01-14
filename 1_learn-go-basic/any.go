package main

import "fmt"

func Ups() any { // any or interface{}
	//return 1
	//return true
	return "Ups"
}
func main() {
	var empty any = Ups()
	fmt.Println(empty)
}
