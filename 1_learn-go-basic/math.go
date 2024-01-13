package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	var c = a + b
	var d = a - b
	var e = a / b
	var f = a * b
	var g = a % b

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	/* augmented assignments */
	var h = 15
	h += 15 // h = h + 15
	fmt.Println(h)
	h -= 20 // h = h - 20
	fmt.Println(h)

	/* unary operator */
	var i = 20
	i++
	fmt.Println(i)
	i--
	fmt.Println(i)
}
