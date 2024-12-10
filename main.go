package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var task string

type requestBody struct {
	Message string `json:"message"` // Структура для JSON
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, task")
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var request requestBody

	// Декодируем JSON из тела запроса
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Обновляем глобальную переменную
	task = request.Message

	// Отправляем ответ
	fmt.Fprintln(w, "Task received:", task)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/get", GetHandler).Methods("GET")

	router.HandleFunc("/post", PostHandler).Methods("POST")

	http.ListenAndServe(":8080", router)
}
