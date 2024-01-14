package main

import (
	"errors"
	"fmt"
)

func Division(value int, divider int) (int, error) {
	if divider == 0 {
		return 0, errors.New("divide by 0")
	} else {
		return value / divider, nil
	}
}

func main() {
	result, err := Division(100, 0)
	if err == nil {
		fmt.Println("result = ", result)
	} else {
		fmt.Println("Error:", err.Error())
	}
}
