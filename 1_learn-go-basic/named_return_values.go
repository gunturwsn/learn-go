package main

import "fmt"

func getCompletedName() (firstName, middleName, lastName string) {
	firstName = "Budi"
	middleName = "Yaya"
	lastName = "Hari"

	return firstName, middleName, lastName
}

func main() {
	firstName, middleName, lastName := getCompletedName()
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}
