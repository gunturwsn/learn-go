package main

import "fmt"

func endApp() {
	fmt.Println("End App")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("ERROR")
	}
}

func main() {
	/*
		Panic function is a function that we can use to stop the program.
		The panic function is usually called when a panic occurs while our program is running.
		When the panic function is called, the program will stop, but the defer function will still be executed
	*/
	runApp(false)
	runApp(true)
}
