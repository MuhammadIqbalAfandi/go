package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func SayHello(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")

	if name == "" {
		fmt.Fprint(writer, "Hello")
	} else {
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestQueryParams(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/hello?name=Cia", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func MultipleParameter(writer http.ResponseWriter, request *http.Request) {
	firstName := request.URL.Query().Get("first_name")
	lastName := request.URL.Query().Get("last_name")
	fmt.Fprintf(writer, "%s %s", firstName, lastName)
}

func TestMultipleParameter(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/hello?first_name=Cia&last_name=Ciaa", nil)
	recorder := httptest.NewRecorder()

	MultipleParameter(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func MultipleParameterSameValue(writer http.ResponseWriter, request *http.Request) {
	var query url.Values = request.URL.Query()
	var names []string = query["name"]
	fmt.Fprint(writer, strings.Join(names, ", "))
}

func TestMultipleParameterSameValue(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/hello?name=Cia&name=Ciaa", nil)
	recorder := httptest.NewRecorder()

	MultipleParameterSameValue(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
