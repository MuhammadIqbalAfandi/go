package goweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SimpleHTMLFileEmbed(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/simple.gohtml"))
	t.ExecuteTemplate(writer, "simple.gohtml", "Hello My name is Ciaa")
}

func TestSimpleHTMLFileEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLFileEmbed(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TemplateDiractory(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseGlob("./templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(writer, "simple.gohtml", "Hello My name is Ciaa")
}

func TestTemplateDiractory(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateDiractory(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
