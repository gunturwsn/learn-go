package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Semarang", "Central Java", "Indonesia"}
	var address2 *Address = &address1
	address2.City = "Magelang"
	fmt.Println(address1)
	fmt.Println(address2)

	/* create new reference */
	//address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	//fmt.Println(address1)
	//fmt.Println(address2)

	/* change by pointer */
	*address2 = Address{"Bandung", "West Java", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)

	var address3 *Address = &address1
	fmt.Println(address3)

}
