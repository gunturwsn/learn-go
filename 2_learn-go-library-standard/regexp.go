package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`b([a-z][a-z])i`)

	fmt.Println(regex.MatchString("budi"))
	fmt.Println(regex.MatchString("bani"))
	fmt.Println(regex.MatchString("buDi"))

	fmt.Println(regex.FindAllString("budi bayu bani buni buDi bAni", 10))
}
