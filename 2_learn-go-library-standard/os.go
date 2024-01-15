package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	for _, arg := range args {
		fmt.Println(arg)
	}

	// running => go run os.go study golang
	// running => go run os.go hello golang, nice to learn you
	// running => go run os.go "hello golang, nice to learn you" (with double quotation marks)

	hostName, err := os.Hostname()
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println(hostName)
	}

}
