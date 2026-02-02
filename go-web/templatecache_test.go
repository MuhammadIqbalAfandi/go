package go_web

import (
	"embed"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed templates/*.gohtml templates/partials/*.gohtml
var templateCache embed.FS

var myTemplates = template.Must(template.ParseFS(templateCache, "templates/*.gohtml", "templates/partials/*.gohtml"))

func TemplateCaching(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "ciaatemplate", map[string]any {
		"Name": "Ciaa",
	})
}

func TestTemplateCaching(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateCaching(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	println(string(body))
}