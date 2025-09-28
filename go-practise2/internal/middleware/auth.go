package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
)


func CheckAuth(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		
		fmt.Println(request.Method, request.URL.Path)

		writer.Header().Set("Content-Type", "application/json")

		apiKey := request.Header.Get("x-api-key")
		if apiKey != "secret123" {
			writer.WriteHeader(http.StatusUnauthorized)
			err_message := map[string]string{"error": "unauthorized"}
			json.NewEncoder(writer).Encode(err_message)
			return
		}

		f(writer, request)
	}
}