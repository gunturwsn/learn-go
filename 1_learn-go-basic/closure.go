package main

import "fmt"

func main() {
	counter := 0

	// function can access data outside the scope
	increment := func() {
		fmt.Println("Increment")
		counter++
	}
	increment()
	increment()
	fmt.Println(counter)
}
