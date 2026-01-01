//go:build ignore

package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "Iqbal, 27\n" +
		"adi, 35\n" +
		"joko, 20"

		reader := csv.NewReader(strings.NewReader(csvString))

		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			}

			fmt.Println(record)
		}
}