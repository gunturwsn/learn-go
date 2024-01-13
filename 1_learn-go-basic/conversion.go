package main

import "fmt"

func main() {
	var value32 int32 = 32768
	var value64 int64 = int64(value32)
	var value16 int16 = int16(value32)

	fmt.Println(value32) // 32768
	fmt.Println(value64) // 32768
	fmt.Println(value16) // -32768 -> Number Overflow
	/* Number Overflow -> The value will return to the lower bound if it exceeds the upper bound */

	var name = "Budi"
	var e = name[0]
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)

}
