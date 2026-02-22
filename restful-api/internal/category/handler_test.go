package category

import (
	"database/sql"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"restfull-api/internal/config"
	"restfull-api/internal/database"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func setupDb() *sql.DB {
	os.Setenv("APP_ENV", "testing")

	cfg := config.Load()
	db, err := database.NewPostgres(cfg.Database)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func setUpRouter() http.Handler {
	router := httprouter.New()

	db := setupDb()
	validate := validator.New()

	categoryRepository := NewPostgres()
	categoryService := NewService(categoryRepository, db, validate)
	categoryHandler := NewHandler(categoryService)
	categoryRoute := NewRoute(categoryHandler)
	categoryRoute.RegisterRoutes(router)

	return router
}

func TestCreateCategorySuccess(t *testing.T) {
	router := setUpRouter()

	requestBody := strings.NewReader(`{"name": "shoes"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestCreateCategoryFail(t *testing.T) {
	router := setUpRouter()

	requestBody := strings.NewReader(`{"name": ""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusInternalServerError, response.StatusCode)
}
