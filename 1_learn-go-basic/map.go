package main

import "fmt"

func main() {
	//var person map[string]string = map[string]string{}
	person := map[string]string{
		"name":    "Budi",
		"address": "Jakarta",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	fmt.Println(person["Birthday"]) // if key doesn't exists, it will be return default value

	/* function for map */
	fmt.Println(len(person))

	person["name"] = "Hari"
	fmt.Println(person)

	delete(person, "address")
	fmt.Println(person)
}
