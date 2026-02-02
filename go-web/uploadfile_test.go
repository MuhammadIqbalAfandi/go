package go_web

import (
	"bytes"
	_ "embed"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

//go:embed resources/profile.jpg
var logo []byte

func UploadFile(writer http.ResponseWriter, request *http.Request) {
	err := myTemplates.ExecuteTemplate(writer, "upload.form.gohtml", nil)
	if err != nil {
		panic(err)
	}
}

func Upload(writer http.ResponseWriter, request *http.Request) {
	file, fileHeader, err := request.FormFile("file")
	if err != nil {
		panic(err)
	}
	fileDestination, err := os.Create("./resources/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}
	name := request.PostFormValue("name")
	myTemplates.ExecuteTemplate(writer, "upload.success.gohtml", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
}

func TestServeUploadFile(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadFile)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestUploadFile(t *testing.T) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("name", "Ciaa")

	file, _ := writer.CreateFormFile("file", "profile.jpg")
	file.Write(logo)
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/upload", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	Upload(recorder, request)

	response := recorder.Result()
	data, _ := io.ReadAll(response.Body)
	println(string(data))
}
