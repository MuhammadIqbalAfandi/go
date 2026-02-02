package go_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateIf(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/if.gohtml"))

	t.ExecuteTemplate(writer, "if.gohtml", map[string]any {
		"Name" : "Ciaa",
	})
}

func TestTemplateIf(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateIf(recorder, request)
	
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TemplateComparator(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/comparator.gohtml"))

	t.ExecuteTemplate(writer, "comparator.gohtml", map[string]any {
		"Name" : "Biaa",
	})
}

func TestTemplateComparator(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateComparator(recorder, request)
	
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TemplateRange(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))

	t.ExecuteTemplate(writer, "range.gohtml", map[string]any {
		"Hobbies": []string {
			"Gaming",
			"Traveling",
			"Cooking",
			"Reading",
		},
	})
}

func TestTemplateRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateRange(recorder, request)
	
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TemplateWith(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/with.gohtml"))

	t.ExecuteTemplate(writer, "with.gohtml", map[string]any {
		"Name": "Ciaa",
		"Address": map[string]any {
			"Street": "123 Main Street",
			"City": "New York",
			"ZipCode": "10001",
		}	,
	})
}

func TestTemplateWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateWith(recorder, request)
	
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TemplateLayout(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/partials/layout.gohtml", "./templates/partials/header.gohtml", "./templates/partials/footer.gohtml"))

	t.ExecuteTemplate(writer, "ciaatemplate", map[string]any {
		"Title": "Welcome to My Website",
		"Name": "Ciaa",
		"Address": map[string]any {
			"Street": "123 Main Street",
			"City": "New York",
			"ZipCode": "10001",
		}	,
	})
}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)
	
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
