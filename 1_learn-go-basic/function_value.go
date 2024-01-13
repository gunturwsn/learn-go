package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	goodby := getGoodBye
	fmt.Println(goodby("Budi"))
}
