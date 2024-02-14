package learn_go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJSON(t *testing.T) {
	jsonRequest := `{"FirstName":"Budi","MiddleName":"Hari","LastName":"Yaya","Age":28,"Married":true}`
	jsonBytes := []byte(jsonRequest)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
}
