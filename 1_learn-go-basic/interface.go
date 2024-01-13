package main

import "fmt"

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

type Bio struct {
	Name string
	Age  int
}

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

func (bio Bio) GetName() string {
	return bio.Name
}

func main() {
	person := Person{Name: "Budi"}
	SayHello(person)

	animal := Animal{Name: "Cat"}
	SayHello(animal)

	bio := Bio{Name: "Hari", Age: 18}
	SayHello(bio)
}
