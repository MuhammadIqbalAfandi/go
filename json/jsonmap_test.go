package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonMap(t *testing.T) {
	jsonRequest := `{"id":"P001","name":"Laptop","price":15000000}`
	jsonBytes := []byte(jsonRequest)

	var result map[string]any
	_ = json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
}
