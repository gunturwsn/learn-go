package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"John", "Paul", "George", "Ringo"}
	values := []int{100, 95, 80, 90}

	fmt.Println(slices.Min(values))
	fmt.Println(slices.Min(names))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Contains(values, 70))
	fmt.Println(slices.Contains(names, "Paul"))
	fmt.Println(slices.Index(values, 1))
	fmt.Println(slices.Index(names, "George"))
}
