package main

import "fmt"

func endApps() {
	fmt.Println("End App")
	message := recover() // correct recover program code
	fmt.Println("An error occurred", message)
}

// wrong recover program code
func runApps2(error bool) {
	defer endApps()
	if error {
		panic("ERROR")
	}
	message := recover()
	fmt.Println("An error occurred", message)
}

// correct recover program code
func runApps(error bool) {
	defer endApps()
	if error {
		panic("ERROR")
	}
}

func main() {
	/*
		Recover is a function that we can use to catch panic.
		With recover the process will stop, so the program will continue to run.
		Put recover inside defer function
	*/
	runApps(true)
	fmt.Println("Run program")
}
