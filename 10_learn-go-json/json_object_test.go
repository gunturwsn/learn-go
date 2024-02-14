package learn_go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Budi",
		MiddleName: "Hari",
		LastName:   "Yaya",
		Age:        28,
		Married:    true,
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
