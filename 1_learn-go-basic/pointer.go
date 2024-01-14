package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	/* pass by value (not using pointer) */
	address1 := Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	address2 := address1
	address2.City = "Bandung"
	fmt.Println(address1) // not change
	fmt.Println(address2) // the city is change

	/* pass by reference using pointer */
	var address3 Address = Address{"Semarang", "Central Java", "Indonesia"}
	var address4 *Address = &address3
	address4.City = "Magelang"
	fmt.Println(address3)
	fmt.Println(address4)

	var data1 int = 5
	var data2 *int = &data1

	fmt.Println(data1)
	fmt.Println(&data1)
	fmt.Println(data2)
	fmt.Println(*data2)
}
