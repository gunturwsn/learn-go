package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()
	data.PushBack("Budi")
	data.PushBack("Hari")
	data.PushBack("Yaya")

	var head *list.Element = data.Front()
	var tail *list.Element = data.Back()
	var next = head.Next()
	fmt.Println(data)
	fmt.Println(head.Value)
	fmt.Println(tail.Value)
	fmt.Println(next.Value)

	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
}
