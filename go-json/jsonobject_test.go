package gojson

import (
	"encoding/json"
	"testing"
)

type Customer struct {
	FirstName string
	LastName  string
	Age       int
	Email     string
	Hobbies   []string
}

func TestJsonObject(t *testing.T) {
	customer := Customer{
		FirstName: "Cia",
		LastName:  "Ciaa",
		Age:       23,
		Email:     "ciaciaa@gmail.com",
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		t.Error(err)
	}
	println(string(bytes))
}
