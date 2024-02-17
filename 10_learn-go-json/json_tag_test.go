package learn_go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func TestJSONTagEncode(t *testing.T) {
	product := Product{
		Id:       "P0001",
		Name:     "Macbook Pro M1 Pro",
		ImageURL: "http://example.com/image.png",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"id":"P0001","name":"Macbook Pro M1 Pro","image_url":"http://example.com/image.png"}`
	jsonByte := []byte(jsonString)

	product := &Product{}
	err := json.Unmarshal(jsonByte, product)
	if err != nil {
		panic(err)
	}
	fmt.Println(product)
}
