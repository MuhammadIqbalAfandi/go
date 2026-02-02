package go_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateString(writer http.ResponseWriter, request *http.Request) {
	templateText := `<html><body><h1>{{.}}</h1></body></html>`
	// t, err := template.New("template-string").Parse(templateText)
	// if err != nil {
	// 	panic(err)
	// }

	t := template.Must(template.New("template-string").Parse(templateText))

	t.ExecuteTemplate(writer, "template-string", "Hello my name is Ciaa")
}

func TestTemplateString(t *testing.T) {
	request := httptest.NewRequest("GET", "http:localhost/", nil)
	recorder := httptest.NewRecorder()

	TemplateString(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}