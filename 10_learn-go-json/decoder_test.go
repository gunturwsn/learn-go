package learn_go_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamDecoder(t *testing.T) {
	reader, _ := os.Open("Customer.json")
	decoder := json.NewDecoder(reader)

	customer := &Customer{}
	err := decoder.Decode(customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
}
