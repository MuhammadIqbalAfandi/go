package main

import (
	"log"
	"net/http"
	"restfull-api/internal/app/router"
	"restfull-api/internal/shared"

	"github.com/joho/godotenv"
)

// LoggingMiddleware logs incoming requests with server port
func LoggingMiddleware(port string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("[%s] %s Port: %s", r.Method, r.RequestURI, port)
			next.ServeHTTP(w, r)
		})
	}
}

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Run server
	host := shared.GetEnv("APP_HOST", "localhost")
	port := shared.GetEnv("APP_PORT", "")

	// Debug: tampilkan nilai environment variable
	log.Printf("DEBUG - APP_HOST: '%s', APP_PORT: '%s'", host, port)

	// Log server startup with port
	log.Printf("Starting server on %s:%s", host, port)

	// Wrap router with logging middleware
	loggedHandler := LoggingMiddleware(port)(router.NewRouter())

	server := http.Server{
		Addr:    host + ":" + port,
		Handler: http.StripPrefix("/api", loggedHandler),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
