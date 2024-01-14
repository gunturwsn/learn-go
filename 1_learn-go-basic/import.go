package main

import (
	"fmt"
	"lear-go-basic/helper"
)

func main() {
	/* package & import */
	budi := helper.SayHello("Budi")
	fmt.Println(budi)

	/* access modifier */
	// helper.version -> cannot access
	fmt.Println(helper.Application)
	//helper.sayGoodBye("Budi") -> cannot access
}
