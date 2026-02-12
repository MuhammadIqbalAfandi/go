package main

import (
	"log"
	"net/http"
	"restfull-api/internal/category"
	"restfull-api/internal/config"
	"restfull-api/internal/infra/database"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {
	// Database config
	cfg := config.Load()
	db, err := database.NewPostgres(cfg.Database)
	if err != nil {
		log.Fatal(err)
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	// Category
	categoryRepository := category.NewPostgres()
	categoryService := category.NewService(categoryRepository, db, validate)
	categoryHandler := category.NewHandler(categoryService)

	// Routing
	router := httprouter.New()

	router.GET("/api/categories", categoryHandler.FindAll)
	router.GET("/api/categories/:categoryId", categoryHandler.FindById)
	router.POST("/api/categories", categoryHandler.Create)
	router.PUT("/api/categories/:categoryId", categoryHandler.Update)
	router.DELETE("/api/categories/:categoryId", categoryHandler.Delete)

	// Run server
	server := http.Server{
		Addr:    "localhost:" + cfg.AppPort,
		Handler: router,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
