package main

import "fmt"

func main() {
	/* variable with init data type */
	var city string
	city = "South Jakarta"
	fmt.Println(city)
	city = "West Jakarta"
	fmt.Println(city)

	/* variable without init data type */
	var cityName = "South Jakarta"
	fmt.Println(cityName)

	/* variable without 'var' keyword */
	place := "Central jakarta"
	fmt.Println(place)

	/* multiple variable */
	var (
		firstName = "Budi"
		lastName  = "Hari"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)

}
