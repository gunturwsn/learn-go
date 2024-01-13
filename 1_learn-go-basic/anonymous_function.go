package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked")
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "Hari"
	}
	registerUser("Hari", blacklist)
	registerUser("Budi", func(name string) bool {
		return name == "Hari"
	})
}
