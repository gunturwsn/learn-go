package learn_go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArrayEncode(t *testing.T) {
	customer := Customer{
		FirstName:  "Budi",
		MiddleName: "Hari",
		LastName:   "Yaya",
		Hobbies:    []string{"Running", "Reading", "Coding"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Budi","MiddleName":"Hari","LastName":"Yaya","Age":0,"Married":false,"Hobbies":["Running","Reading","Coding"]}`
	jsonByte := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonByte, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplexEncode(t *testing.T) {
	customer := Customer{
		FirstName: "Budi",
		Addresses: []Address{
			{
				Street:     "Jalan Kenanga",
				Country:    "Indonesia",
				PostalCode: "12233",
			},
			{
				Street:     "Jalan Flamboyan",
				Country:    "Indonesia",
				PostalCode: "13344",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Budi","MiddleName":"","LastName":"","Age":0,"Married":false,"Hobbies":null,"Addresses":[{"Street":"Jalan Kenanga","Country":"Indonesia","PostalCode":"12233"},{"Street":"Jalan Flamboyan","Country":"Indonesia","PostalCode":"13344"}]}`
	jsonByte := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonByte, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addresses)
}

func TestOnlyJSONArrayComplexEncode(t *testing.T) {
	address := []Address{
		{
			Street:     "Jalan Kenanga",
			Country:    "Indonesia",
			PostalCode: "12233",
		},
		{
			Street:     "Jalan Flamboyan",
			Country:    "Indonesia",
			PostalCode: "13344",
		},
	}

	bytes, _ := json.Marshal(address)
	fmt.Println(string(bytes))
}

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonArray := `[{"Street":"Jalan Kenanga","Country":"Indonesia","PostalCode":"12233"},{"Street":"Jalan Flamboyan","Country":"Indonesia","PostalCode":"13344"}]`
	jsonByte := []byte(jsonArray)

	address := &[]Address{}
	err := json.Unmarshal(jsonByte, address)
	if err != nil {
		panic(err)
	}
	fmt.Println(address)
}
