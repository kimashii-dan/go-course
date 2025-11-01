package main

import (
	"fmt"
	"go-practise5/internal/db"
	"go-practise5/internal/handlers"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.Connect()

	router := http.NewServeMux()

	router.HandleFunc("GET /movies", handlers.GetMovies)
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Server starting error: ", err)
	}
}
