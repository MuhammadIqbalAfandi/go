package router

import (
	"log"
	"restfull-api/internal/category"
	"restfull-api/internal/config"
	"restfull-api/internal/database"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func NewRouter() *httprouter.Router {
	router := httprouter.New()

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
	categoryRoute := category.NewRoute(categoryHandler)
	categoryRoute.RegisterRoutes(router)

	return router
}
