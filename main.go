package main

import (
	"encoding/json"
	"net/http"
	"strconv"

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

// PATCH: обновить задачу по ID
func PatchHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем ID из параметров URL
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Находим задачу по ID
	var task Task
	if err := DB.First(&task, id).Error; err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	// Декодируем JSON из тела запроса
	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Обновляем задачу
	if err := DB.Model(&task).Updates(updates).Error; err != nil {
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}

	// Возвращаем обновленную задачу
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// DELETE: удалить задачу по ID
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем ID из параметров URL
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Находим задачу по ID
	var task Task
	if err := DB.First(&task, id).Error; err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	// Удаляем задачу
	if err := DB.Delete(&task).Error; err != nil {
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}

	// Возвращаем успех
	w.WriteHeader(http.StatusNoContent)
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
	router.HandleFunc("/patch/{id}", PatchHandler).Methods("PATCH")    // Новый маршрут
	router.HandleFunc("/delete/{id}", DeleteHandler).Methods("DELETE") // Новый маршрут

	// Запускаем сервер
	http.ListenAndServe("localhost:8080", router)
}
