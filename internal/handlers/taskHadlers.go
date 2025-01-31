package handlers

import (
	"context"
	"moy_proekt/internal/taskService"
	"moy_proekt/internal/web/tasks"
)

type TaskHandler struct {
	service *taskService.TaskService
}

func NewTaskHandler(service *taskService.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

// Получение всех задач
func (h *TaskHandler) GetTasks(ctx context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	tasksList, err := h.service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}
	for _, task := range tasksList {
		response = append(response, tasks.Task{
			Id:     toIntPtr(int(task.ID)),
			Task:   toStringPtr(task.Task),
			IsDone: toBoolPtr(task.IsDone),
			UserId: toIntPtr(int(task.UserID)),
		})
	}
	return &response, nil
}

// Получение всех задач конкретного пользователя
func (h *TaskHandler) GetUsersUserIdTasks(ctx context.Context, req tasks.GetUsersUserIdTasksRequestObject) (tasks.GetUsersUserIdTasksResponseObject, error) {
	tasksList, err := h.service.GetTasksByUserID(uint(req.UserId))
	if err != nil {
		return nil, err
	}

	response := tasks.GetUsersUserIdTasks200JSONResponse{}
	for _, task := range tasksList {
		response = append(response, tasks.Task{
			Id:     toIntPtr(int(task.ID)),
			Task:   toStringPtr(task.Task),
			IsDone: toBoolPtr(task.IsDone),
			UserId: toIntPtr(int(task.UserID)),
		})
	}
	return &response, nil
}

// Создание задачи
func (h *TaskHandler) PostTasks(ctx context.Context, req tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	newTask := taskService.Task{
		Task: req.Body.Task,
	}

	if req.Body.IsDone != nil {
		newTask.IsDone = *req.Body.IsDone
	}

	if req.Body.UserId != nil {
		newTask.UserID = uint(*req.Body.UserId)
	}

	createdTask, err := h.service.CreateTask(newTask)
	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:     toIntPtr(int(createdTask.ID)),
		Task:   toStringPtr(createdTask.Task),
		IsDone: toBoolPtr(createdTask.IsDone),
		UserId: toIntPtr(int(createdTask.UserID)),
	}
	return &response, nil
}

// Обновление задачи
func (h *TaskHandler) PatchTasksId(ctx context.Context, req tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	updatedTask := taskService.Task{}

	if req.Body.Task != nil {
		updatedTask.Task = *req.Body.Task
	}
	if req.Body.IsDone != nil {
		updatedTask.IsDone = *req.Body.IsDone
	}
	if req.Body.UserId != nil {
		updatedTask.UserID = uint(*req.Body.UserId)
	}

	result, err := h.service.UpdateTaskByID(uint(req.Id), updatedTask)
	if err != nil {
		return nil, err
	}

	response := tasks.PatchTasksId200JSONResponse{
		Id:     toIntPtr(int(result.ID)),
		Task:   toStringPtr(result.Task),
		IsDone: toBoolPtr(result.IsDone),
		UserId: toIntPtr(int(result.UserID)),
	}
	return &response, nil
}

// Удаление задачи
func (h *TaskHandler) DeleteTasksId(ctx context.Context, req tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	if err := h.service.DeleteTaskByID(uint(req.Id)); err != nil {
		return nil, err
	}
	return &tasks.DeleteTasksId204Response{}, nil
}

var _ tasks.GetTasksIdResponseObject = (*tasks.GetTasksId404Response)(nil)

// Получение задачи по ID
func (h *TaskHandler) GetTasksId(ctx context.Context, req tasks.GetTasksIdRequestObject) (tasks.GetTasksIdResponseObject, error) {
	task, err := h.service.GetTaskByID(uint(req.Id))
	if err != nil {
		return nil, err
	}

	// Если задача не найдена, возвращаем 404 ответ
	if task == nil {
		return tasks.GetTasksIdResponseObject(tasks.GetTasksId404Response{}), nil
	}

	// Формируем 200 ответ с данными задачи
	response := tasks.GetTasksId200JSONResponse{
		Id:     toIntPtr(int(task.ID)),
		Task:   toStringPtr(task.Task),
		IsDone: toBoolPtr(task.IsDone),
		UserId: toIntPtr(int(task.UserID)),
	}
	return &response, nil // Возвращаем указатель для 200 ответа
}
