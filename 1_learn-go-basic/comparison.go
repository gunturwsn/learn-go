package main

import "fmt"

func main() {
	var name1 = "Budi"
	var name2 = "Budi"

	var result bool = name1 == name2
	fmt.Println(result)

	result = name1 != name2
	fmt.Println(result)

	result = name1 > name2
	fmt.Println(result)

	result = name1 < name2
	fmt.Println(result)
}
