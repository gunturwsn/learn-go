package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Budi"
	names[1] = "Hari"
	names[2] = "Mukti"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90, 80, 95,
	}
	fmt.Println(values)

	/* function for array */
	fmt.Println(len(values))
	fmt.Println(values[1])
	values[1] = 50
	fmt.Println(values[1])

	var datas = [...]int{
		90, 80, 95, 80, 78,
	}
	fmt.Println(datas)
}
