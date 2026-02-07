package gojson

import (
	"encoding/json"
	"os"
	"testing"
)

func TestStreamEncode(t *testing.T) {
	writer, _ := os.Create("sample_output.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName: "Cia",
		LastName:  "Ciaa",
		Age:       23,
		Email:     "ciaciaa@gmail.com",
	}

	_ = encoder.Encode(customer)
}