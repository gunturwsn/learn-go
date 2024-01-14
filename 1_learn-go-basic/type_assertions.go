package main

import "fmt"

func random() any {
	return 100
}

func main() {
	var result any = random()
	//var resultString string = result.(string)
	//fmt.Println(resultString)

	//var resultInt int = result.(int)
	//fmt.Println(resultInt)

	/* type asserions switch for check data type */
	switch value := result.(type) {
	case string:
		fmt.Println("string", value)
	case int:
		fmt.Println("int", value)
	default:
		fmt.Println("unknown")
	}
}
