package goweb

import (
	"fmt"
	"net/http"
	"testing"
)

func DownloadFile(writer http.ResponseWriter, request *http.Request) {
	fileName := request.URL.Query().Get("file")
	if fileName == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "BAD REQUEST")
		return
	}

	writer.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	http.ServeFile(writer, request, "./resources/"+fileName)
}

func TestServeDownloadFile(t *testing.T) {
	serve := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(DownloadFile),
	}

	err := serve.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
