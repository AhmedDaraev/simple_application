package handlers

import (
	"encoding/json"
	"moy_proekt/internal/taskService"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	Service *taskService.TaskService
}

func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.Service.GetAllTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(tasks) == 0 {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "No tasks found"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (h *Handler) PostTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task taskService.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdTask, err := h.Service.CreateTask(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdTask)
}

func (h *Handler) PatchTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Извлекаем ID задачи из URL
	id := mux.Vars(r)["id"]
	taskID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	// Декодируем тело запроса в task
	var task taskService.Task
	err = json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Обновляем задачу по ID
	updatedTask, err := h.Service.UpdateTaskByID(int(taskID), task) // Приводим taskID к uint
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправляем обновленную задачу в ответ
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTask)
}

func (h *Handler) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Извлекаем ID задачи из URL
	id := mux.Vars(r)["id"]
	taskID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	// Удаляем задачу по ID
	err = h.Service.DeleteTaskByID(int(taskID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204 No Content, если удаление прошло успешно
}
