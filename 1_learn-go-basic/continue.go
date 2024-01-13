package main

import "fmt"

func main() {
	for index := 0; index < 10; index++ {
		if index%2 == 0 {
			continue
		}
		fmt.Println("index", index)
	}
}
