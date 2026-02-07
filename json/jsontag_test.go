package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	ImageUrl string `json:"image_url"`
}

func TestJsonTag(t *testing.T) {
	product := Product{
		Id:       "P001",
		Name:     "Laptop",
		Price:    15000000,
		ImageUrl: "http://example.com/laptop.jpg",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
