package main

import (
	"fmt"
	"lear-go-basic/database"
	_ "lear-go-basic/internal"
)

func main() {
	fmt.Println(database.GetDatabase())
}
