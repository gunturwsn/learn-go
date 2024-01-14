package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesiaByValue(address Address) {
	address.Country = "Indonesia"
}

func ChangeCountryToIndonesiaByPointer(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address1 := Address{}
	ChangeCountryToIndonesiaByValue(address1)
	fmt.Println(address1)

	address2 := &Address{}
	ChangeCountryToIndonesiaByPointer(address2)
	fmt.Println(address2)

	address3 := Address{}
	ChangeCountryToIndonesiaByPointer(&address3)
	fmt.Println(address3)
}
