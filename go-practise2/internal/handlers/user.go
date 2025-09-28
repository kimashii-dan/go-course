package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func GetUser(writer http.ResponseWriter, request *http.Request) {

	query := request.URL.Query()
	id := query.Get("id")

	if id == "" {
		writer.WriteHeader(http.StatusBadRequest)
		err_message := map[string]string{"error": "invalid id"}
        json.NewEncoder(writer).Encode(err_message)
        return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)	
		err_message := map[string]string{"error": "invalid id"}
        json.NewEncoder(writer).Encode(err_message)
        return
	}

	// if idInt < 1 {
	// 	writer.WriteHeader(http.StatusBadRequest)	
	// 	err_message := map[string]string{"error": "invalid id"}
    //     json.NewEncoder(writer).Encode(err_message)
    //     return
	// }

	writer.WriteHeader(http.StatusOK)
	response := map[string]int{"user_id": idInt}
	json.NewEncoder(writer).Encode(response)
}

type User struct {
	Name string
}

func CreateUser(writer http.ResponseWriter, request *http.Request){

	var u User
	err := json.NewDecoder(request.Body).Decode(&u)

	writer.Header().Set("Content-Type", "application/json")
	
	if err != nil {
        writer.WriteHeader(http.StatusBadRequest)
		err_message := map[string]string{"error": "invalid name"}
		json.NewEncoder(writer).Encode(err_message)
        return
	}

	if u.Name == "" {
        writer.WriteHeader(http.StatusBadRequest)
		err_message := map[string]string{"error": "invalid name"}
        json.NewEncoder(writer).Encode(err_message)
        return
	}

    writer.WriteHeader(http.StatusCreated)
	response := map[string]string{"created": u.Name}
	json.NewEncoder(writer).Encode(response)
}