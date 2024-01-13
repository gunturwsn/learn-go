package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	total := sumAll(10, 10, 10, 10, 10)
	fmt.Println(total)

	/* if we already have data in slice */
	numbers := []int{10, 10, 10, 10, 10, 10, 10, 10, 10}
	total = sumAll(numbers...)
	fmt.Println(total)
}
