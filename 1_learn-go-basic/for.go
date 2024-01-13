package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("counter =", counter)
		counter++
	}

	/* for with statement */
	/* for has 2 statement:
	-> init statement = statements before executed
	-> post statement = statements are always executed at the end
	*/
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("counter =>", counter)
	}

	/* for range */
	names := []string{"Budi", "Hari", "Yaya"}
	for index, value := range names {
		fmt.Println("Index", index, "Hello", value)
	}

	for _, value := range names {
		fmt.Println("Hello", value)
	}
}
