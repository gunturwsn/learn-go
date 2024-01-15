package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Budi Hari", "di"))
	fmt.Println(strings.Split("Budi Hari", " "))
	fmt.Println(strings.ToLower("Budi Hari"))
	fmt.Println(strings.ToUpper("Budi Hari"))
	fmt.Println(strings.Trim("     Budi Hari     ", " "))
	fmt.Println(strings.ReplaceAll("Budi Hari", "di", "at"))
}
