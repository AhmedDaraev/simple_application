package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []Task

	// Получаем все записи из таблицы
	if err := DB.Find(&tasks).Error; err != nil {
		http.Error(w, "Failed to retrieve tasks", http.StatusInternalServerError)
		return
	}

	// Возвращаем слайс задач в JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var taskInput Task

	// Декодируем JSON из тела запроса
	if err := json.NewDecoder(r.Body).Decode(&taskInput); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Сохраняем задачу в БД
	if err := DB.Create(&taskInput).Error; err != nil {
		http.Error(w, "Failed to save task to database", http.StatusInternalServerError)
		return
	}

	// Возвращаем сохранённую задачу
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(taskInput)
}

func main() {
	// Инициализация БД
	InitDB()

	// Автоматическая миграция модели Task
	DB.AutoMigrate(&Task{})

	// Создаём маршруты
	router := mux.NewRouter()
	router.HandleFunc("/get", GetHandler).Methods("GET")
	router.HandleFunc("/post", PostHandler).Methods("POST")

	// Запускаем сервер
	http.ListenAndServe("localhost:8080", router)
}
