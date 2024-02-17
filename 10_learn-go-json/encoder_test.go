package learn_go_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamEncoder(t *testing.T) {
	writter, _ := os.Create("Customer_output.json")
	encoder := json.NewEncoder(writter)

	customer := Customer{
		FirstName:  "Budi",
		MiddleName: "Hari",
		LastName:   "Yaya",
		Age:        28,
		Married:    false,
		Hobbies: []string{
			"Reading",
			"Coding",
		},
		Addresses: []Address{
			{
				Street:     "Jalan Panjang",
				Country:    "Indonesia",
				PostalCode: "123345",
			},
		},
	}

	err := encoder.Encode(customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
}
