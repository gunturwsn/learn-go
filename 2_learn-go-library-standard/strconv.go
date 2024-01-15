package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(boolean)
	}

	resultInt, err := strconv.Atoi("31")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(resultInt)
	}

	binary := strconv.FormatInt(999, 2)
	fmt.Println(binary)

	var intToString string = strconv.Itoa(888)
	fmt.Println(intToString)
}
