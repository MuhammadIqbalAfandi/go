package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(writer http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("content-type")
	fmt.Fprint(writer, contentType)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost/", nil)
	request.Header.Add("Content-Type", "aplication/json")
	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func ResponseHeader(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("X-Powered-By", "CodeFalah")
	fmt.Fprint(writer, "OK")
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost/", nil)
	request.Header.Add("Content-Type", "aplication/json")
	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)

	poweredBy := recorder.Header().Get("X-Powered-By")
	fmt.Println(poweredBy)
}
