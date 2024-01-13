package main

import "fmt"

func main() {
	type NoKtp string

	var ktpBudi NoKtp = "11111111"
	fmt.Println(ktpBudi)
	fmt.Println(NoKtp("222222"))
}
