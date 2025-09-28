package main

import (
	"fmt"
	"net/http"
	"practise2/internal/handlers"
	"practise2/internal/middleware"
)

func main(){
	router := http.NewServeMux()

	router.HandleFunc("GET /user", middleware.CheckAuth(handlers.GetUser))
	router.HandleFunc("POST /user", middleware.CheckAuth(handlers.CreateUser))

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Server starting error: ", err)
	}
}