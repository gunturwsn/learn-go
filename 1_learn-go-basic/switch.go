package main

import "fmt"

func main() {
	name := "Budi"

	switch name {
	case "Hari":
		fmt.Println("Hello Hari")
	case "Budi":
		fmt.Println("Hello Budi")
	default:
		fmt.Println("Hi, nice to meet you")
	}

	/* switch short statement */
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("name is too long")
	case false:
		fmt.Println("proper name")
	}

	/* switch without condition */
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("name is too long")
	case length > 5:
		fmt.Println("name is quite long")
	default:
		fmt.Println("proper name")
	}
}
