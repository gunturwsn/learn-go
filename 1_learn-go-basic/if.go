package main

import "fmt"

func main() {
	name := "Hari"

	if name == "Budi" {
		fmt.Println("Hello Budi")
	} else if name == "Hari" {
		fmt.Println("Hello Hari")
	} else {
		fmt.Println("Hi, nice to meet you")
	}

	/* if short statement */
	if length := len(name); length > 5 {
		fmt.Println("name is too long")
	} else {
		fmt.Println("proper name")
	}
}
