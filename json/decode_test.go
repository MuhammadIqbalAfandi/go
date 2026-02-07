package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {
	jsonObject := `{"FirstName":"Cia","LastName":"Ciaa","Age":23,"Email":"ciaciaa@gmail.com"}`
	jsonBytes := []byte(jsonObject)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		t.Errorf("Error decoding JSON: %v", err)
	}

	fmt.Println(customer.FirstName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
	fmt.Println(customer.Email)
}
