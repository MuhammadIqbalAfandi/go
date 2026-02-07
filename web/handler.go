//go:build ignore

package main

import (
	"fmt"
	"net/http"
)

func main () {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		// logic here
		fmt.Fprint(writer, "Hello World")
	}

	server := http.Server {
		Addr: "localhost:8080",
		Handler: handler,
	}

	server.ListenAndServe()
}