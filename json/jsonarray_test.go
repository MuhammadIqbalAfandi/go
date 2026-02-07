package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonArray(t *testing.T) {
	customer := Customer{
		FirstName: "Cia",
		LastName:  "Ciaa",
		Age:       23,
		Email:     "ciaciaa@gmail.com",
		Hobbies:   []string{"Reading", "Traveling", "Gaming"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
