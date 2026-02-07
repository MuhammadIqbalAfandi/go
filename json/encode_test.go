package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Encode(data any) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestEncode(t *testing.T) {
	Encode("ciaa")
	Encode(123)
	Encode(true)
	Encode([]int{1, 2, 3})
	Encode(map[string]int{"a": 1, "b": 2})
}
