package learn_go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapDecode(t *testing.T) {
	jsonString := `{"id":"P0001","name":"Macbook Pro M1 Pro","image_url":"http://example.com/image.png"}`
	jsonByte := []byte(jsonString)

	var result map[string]interface{}
	err := json.Unmarshal(jsonByte, &result)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["image_url"])
}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":        "P0001",
		"name":      "Macbook Pro M1 Pro",
		"image_url": "http://example.com/image.png",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
