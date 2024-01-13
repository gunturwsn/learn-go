package main

import "fmt"

func logging() {
	fmt.Println("Finished calling the function")
}

func runApplication() {
	defer logging()
	fmt.Println("Run Application")
}

func main() {
	/*
		The defer function is function that we can schedule to be executed after a function has finished executing.
		The defer function will always be executed even if an error occurs in the function being executed.
	*/
	runApplication()
}
