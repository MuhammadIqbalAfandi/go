package goweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateData(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/data.gohtml"))

	t.ExecuteTemplate(writer, "data.gohtml", map[string]interface{}{
		"Title": "Learning Golang Web Template",
		"Name":  "Ciaa",
	})
}

func TestTemplateData(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateData(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

type Page struct {
	Title string
	Name  string
}

func TemplateDataStruct(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/data.gohtml"))

	t.ExecuteTemplate(writer, "data.gohtml", Page{
		Title: "Learning Golang Web Template",
		Name:  "Ciaa",
	})
}

func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
