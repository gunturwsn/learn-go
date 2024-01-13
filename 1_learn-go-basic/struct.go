package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name+", my name is", customer.Name)
}

func main() {
	var budi Customer

	fmt.Println(budi)
	budi.Name = "Budi"
	budi.Address = "Jakarta"
	budi.Age = 28
	fmt.Println(budi)
	fmt.Println(budi.Name)
	fmt.Println(budi.Address)
	fmt.Println(budi.Age)

	/* struct literal */
	hari := Customer{
		Name:    "Hari",
		Address: "Bandung",
		Age:     30,
	}
	fmt.Println(hari)

	// the order of the fields must be the same
	yaya := Customer{"Yaya", "Bogor", 26}
	fmt.Println(yaya)

	/* struct method */
	jaja := Customer{"Jaja", "Bekasi", 32}
	jaja.sayHello("Budi")
}
