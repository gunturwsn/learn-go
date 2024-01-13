package main

import "fmt"

func getFullName() (string, string) {
	return "Budi", "Hari"
}
func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)

	/* ignore return value */
	name, _ := getFullName()
	fmt.Println(name)
}
