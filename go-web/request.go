//go:build ignore

package main

import (
	"fmt"
	"net/http"
)

func main() {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, request.Method)
		fmt.Fprintln(writer, request.Header)
	}

	server := http.Server {
		Addr: "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}