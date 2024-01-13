package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string {
	if name == "Dog" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Budi", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Dog", filter)
}
